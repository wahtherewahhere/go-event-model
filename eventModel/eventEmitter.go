package eventModel

import (
	//"reflect"
	"log"
)

type Event_mange_function func ( interface{}, interface{}) (err error)

type EventEmitter struct {
	eventList map[string][]Event_mange_function
	owner string
	showTriggerInfo bool
}


func GetEventEmitter() (*EventEmitter) {
	eventModel := new(EventEmitter);
	eventModel.eventList = make(map[string][]Event_mange_function);
	eventModel.showTriggerInfo = true;
	// lifeTime expended;
	return eventModel; 
}

func (instance *EventEmitter) SetOwnerName (name string) {
	instance.owner = name;
}

// Setting whether Show the process of Event or not,
// turn OFF this variable WON'T trun off the Error Message from EventModel !!
func (instance *EventEmitter) SetTriggerInfo (onoff bool) {
	instance.showTriggerInfo = onoff;
}

// Event - signification
func (instance *EventEmitter) On (name string, f Event_mange_function) {
	log.Println("[ Event Model ][event reg.] ", instance.owner ," <-> ", name );
	// add function to specific "Eevent Name" 
	instance.eventList[name] = append( instance.eventList[name], f );
}

// Event - trigger the event 
func (instance *EventEmitter) Trigger (name string, from interface{}, data interface{}) {
	// Console info.
	if instance.showTriggerInfo {
		// show it!
		log.Printf("[ Event Model ] %s ---> %s ,    Data: %v", instance.owner, name, data );
	}
	
	// no such event
	if len(instance.eventList[name]) <= 0 {
		log.Println("[ Event Model ][error] no function service event .....>", name );
		return;
	}
	
	// it is
	for index, _ := range instance.eventList[name] {
		///
		err := instance.eventList[name][index]( from, data ); 
		if err != nil {
			log.Println("[error] something wrong.... Event-->", name, "|", err.Error() );
		}
	}
	
}