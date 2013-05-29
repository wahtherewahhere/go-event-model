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
* `Return` pointer of eventModel.EventEmitter
 
get an instnace to control and manager event  
you can have more than one instance whatever you want  
if your server need, give it an Emitter.  
if your engine need, give it an Emitter.  
All is independent to each others.  
Instance ownes its own system!  
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

In default design, 

* <b>first parameter</b>: let you put the information about who trigge, and

* <b>secondary parameter</b>:  is for storing data from event.

But, in fact, you can put anything you want.

Just remeber using 'reflect' to convert to correct type.

Let's all!

Have Fun!


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
    

## License

(The MIT License)

Copyright (C) 2013 wahtherewahhere@github.com

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
