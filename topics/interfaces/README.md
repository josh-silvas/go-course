## Interfaces
Interfaces allow the Go programming language to achieve a polymorphic behavior.  This (in my opinion) is one of the 
most powerful specs in the language because it allows for a Go (being a static-typed language) to define a 
complex type by it's behavior and not necessarily by the value or type.

#### Some Standards
 - Interface Naming
   - By convention, one-method interfaces are named by the method name plus an -er suffix or similar modification to 
   construct an agent noun: Reader, Writer, Formatter, CloseNotifier etc. 
   ```go 
   type Reader interface { 
       Read()
   }
   ```
   - There are a number of such names and it's productive to honor them and the function names they capture. 
   Read, Write, Close, Flush, String and so on have canonical signatures and meanings. To avoid confusion, don't 
   give your method one of those names unless it has the same signature and meaning. Conversely, if your type 
   implements a method with the same meaning as a method on a well-known type, give it the same name and signature; 
   call your string-converter method String not ToString. 
   
 - When/Where to Define Interfaces
   - Do not define interfaces before they are used: without a realistic example of usage, it is too difficult to see 
   whether an interface is even necessary, let alone what methods it ought to contain.


#### Implementing an Interface
Interfaces are created similar to structs from a syntax perspective. With interfaces, you will populate the interface
with the methods you require to satisfy that interface.  

So what does that mean?!
  * There are a lot of docs out there to go over interfaces and interface semantics, and they typically use words that 
  make them more of a mystery. You can think of an interface as a type that is defined by methods (or behaviors) rather
  than static types. I.E. instead of defining a type that contains a string and int, you could define an interface
  that contains a set of methods that a type needs to have before it can match the interface. 

In the example below, we create an interface called `Alerter`, satisfied with an `Alert()` method that
returns an error.
```go
type Alerter interface {
	Alert() error
}
```

With our new definition, we can create a new function that will take our interface as a variable, and call the
method `Alert()` on it.
```go
func Send(a Alerter) error {
	return a.Alert()
}
```

With this interface, we can now create types and define the required methods for each type which will allow it
to satisfy interface.  Below, we create a new type named `Slack`, because the `Slack` type implements a method called
`Alert()` that returns an error, this type now satisfies the `Alerter` interface. 
```go
type Slack struct {
	string
}

func (s *Slack) Alert() error {
	//
	// Do some logic to initiate an API call 
	// to Slack and send the message to the 
	// specified channel.
	//
	return nil
}
```

So using this code, we can now pass any type that satisfies the `Alerter` interface into the `Send()` function.  
```go
func main() {
	if err := Send(&Slack{"test Slack message"}); err != nil {
		// Handle errors better than me
		panic(err)
	}
}
```

This is really cool because we are able to decouple the functions from the types that they required just by creating
an interface that defines a certain set of behavior of a type.

 - [Go Playground Example](https://play.golang.org/p/6j0kOc3QJAO)
 - [Source Example](https://github.com/josh5276/go-course/blob/master/topics/interfaces/example.go)
 
#### References
 - [Interface Naming - Effective Go](https://golang.org/doc/effective_go.html#interface-names)
 - [Code Review Comments - Interfaces](https://github.com/golang/go/wiki/CodeReviewComments#interfaces)