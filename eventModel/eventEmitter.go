// (The MIT License)
//
// Copyright (C) 2013 wahtherewahhere@github.com
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and 
// associated documentation files (the "Software"), 
// to deal in the Software without restriction, including without limitation the rights to 
// use, copy, modify, merge, publish, distribute, sublicense, 
// and/or sell copies of the Software, 
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be 
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, 
// INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, 
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. 
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, 
// DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, 
// ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
//
//
package eventModel

import (
	"log"
	"runtime"
	"reflect"
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
	// catch error
	defer func() {
		if e := recover(); e != nil {}
	}()
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
		if instance.showTriggerInfo {
			log.Println("[ Event Model ] Send Signal to ", getFunctionName(instance.eventList[name][index]) );
		}
		err := instance.eventList[name][index]( from, data ); 
		if err != nil {
			log.Println("[error] something wrong.... Event-->", name, "|", err.Error() );
		}
	}
}

// private
func getFunctionName(i interface{}) string {
    return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
