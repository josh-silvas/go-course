## Methods
Methods in the Go language are similar to functions except they declare a 
receiver variable (named variable or struct type). 

A good way to think of methods are functions that perform actions tied to a specific set of data. Only 
defined sets of the data in the receiver can access methods.

In the example below, we can see that trying to call the method `FullName` as a function will result in
an error. Likewise, calling the function `Test` as a method will result in an error as well.
```go
import "fmt"

type User struct {
	First string
	Last  string
}

func main() {
	r := User{First: "Josh", Last: "Silvas"}
	
	// Cannot call a method without a receiver
	// prog.go:16:2: undefined: FullName
	FullName()
	
	// Cannot call a function with a receiver
	// prog.go:20:3: r.Test undefined (type User has no field or method Test)
	r.Test()
}

func (u User) FullName() {
	fmt.Printf("%s %s\n", u.First, u.Last)
}

func Test() {
	fmt.Println("Test String")
}
``` 
[Go Playground Example](https://play.golang.org/p/xljS5wYReGE)

The concept of value and pointer semantics still apply when using method receivers. That is, a method receiver
of `(u User)` will pass a copy of the value of `User` into the method, whereas a method receiver of 
`(u *User)` will pass the address value of `User`.  When sharing data across program boundaries, be mindful
of the behavior of receiver value semantics.

Methods initialized with a value receiver can be invoked with both a pointer and a 
value, however, a method initialized with a pointer receiver can only be invoked with
a pointer. 

| Type                       | Examples                          |
|----------------------------|-----------------------------------|
| func (t *Type) Example() { | *type.Example()                   |
| func (t Type) Example() {  | type.Example()<br>*type.Example() |

### Example
 - [Source Example](https://github.com/josh5276/go-course/blob/master/topics/methods/example.go)