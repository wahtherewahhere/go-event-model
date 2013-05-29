go-event-model
================

event-model, written by go-lang,   
work fine with Go 1.1 

## Install

`go get github.com/wahtherewahhere/go-event-model/eventModel`

## Import

`import "github.com/wahtherewahhere/go-event-model/eventModel" `  
or use alias  
`import em "github.com/wahtherewahhere/go-event-model/eventModel"`  

## Usage

### GetEventEmitter() (*eventModel.EventEmitter)
get an instnace to control and manager event  
`eventEmitterInstance := eventModel.GetEventEmitter();`

### On(name string, f eventModel.Event_mange_function)
Binding an event

    eventEmitterInstance.On("hello", handler1)
    eventEmitterInstance.On("world", handler2)
    eventEmitterInstance.On("thank", handler3)
    // bla bla....
    
### Trigger(name string)
Signal event by name  

    eventEmitterInstance.Trigger("hello", nil, nil)
    eventEmitterInstance.Trigger("world", from, nil)
    eventEmitterInstance.Trigger("world", from, nil)
    eventEmitterInstance.Trigger("thank", gift, gift2)
    eventEmitterInstance.Trigger("thank", nil, gift)
    eventEmitterInstance.Trigger("thank", gift_array, gift_array2)
    // bla bla....

## Callback Format

`func handler_name (from interface{}, data interface{}) (err error) {}`

have to 2 parameter to let you do everything!

In default design, first parameter let you put the information about who trigge, 

and second is data from event.

But, in fact, you can put anything you want.

Just remeber using 'reflect' to convert to correct type.

Let's all!

Have Fun!

Hase

## Example

    package main
    
    import (
        "fmt"
        "github.com/wahtherewahhere/go-event-model/eventModel"
    )
    
    func helloHandler1 (from interface{}, data interface{}) (err error) {
        fmt.Println("hello event! #1");
    }
    
    func helloHandler2 (from interface{}, data interface{}) (err error) {
        fmt.Println("hello event! #2");
    }
    
    func main () {
        // GetEventEmitter() return a pointer of `eventModel.EventEmitter`
        eventEmitterInstance := eventModel.GetEventEmitter();
        
        // optional - a embedded debug info will show the owner or empty string in console
        eventEmitterInstance.SetOwnerName("main");
        
        // optional - if you want to turn on or off debug message, just call
        eventEmitterInstance.SetTriggerInfo(false);
        eventEmitterInstance.SetTriggerInfo(true);
        
        // binding event called 'hello' with function called 'helloHandler'
        eventEmitterInstance.On("hello", helloHandler1);
        
        // same event can be binded many function with.
        eventEmitterInstance.On("hello", helloHandler2);
        
        // all function which binding event 'hello' will be called
        eventEmitterInstance.Trigger("hello", nil, nil);
        
        fmt.Println("thanks you for download my package!");
    }
    
    
