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
        eventEmitterInstance.Trigger("hello");
        
        fmt.Println("thanks you for download my package!");
    }
    
    
