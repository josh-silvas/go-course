## Pointers
A pointer is a memory address of a given value that allows functions, methods, etc.. 
a way to access data that is shared across a program boundary. That is, to allow 
functions to share a value with another function so that it can read/write to that value 
even though the value does not exist directly inside its own frame.

##### Syntax
Pointers look very similar to a regular variable, with a `&` or `*` distinguisher
prepended to the variable name. 
 * `&varName`: Pass variable by reference. This is the memory "address of" the variable
 not the actual variable value. 
 * `*varName`: Dereference variable. This action will attempt to pull the variable 
 value from the location that the pointer is pointing to.

Passing data by value makes a copy of the data, manipulates it, and _can_ return that
data back to you. Passing data by pointer allows the accepting function to manipulate the data
in it's original memory location.

+ Go initially allocates a 2048 byte block of memory for each function (frame)
  + Additional functions will get allocated their own frame and data can be shared
  between functions by value or by pointers.
+ Don't use pointers on built-in types unless there is a very specific need for a pointer
+ Do use pointers when it is not safe to copy the original data.

> If the nature of the struct is something that should not be changed, 
like a time, a color or a coordinate, then implement the struct as a 
primitive data value. If the nature of the struct is something that
 can be changed, even if it never is in your program, it is not a 
 primitive data value and should be implemented to be shared with a pointer.

 -- [Using Pointers in Go](https://www.ardanlabs.com/blog/2014/12/using-pointers-in-go.html) 

Take the following example:
```go
func main() {
        var count = 10
    
	fmt.Printf("Count:\t %d | Addr: %v\n", count, &count)
	
	addOneByValue(count)
	fmt.Printf("Count:\t %d | Addr: %v\n", count, &count)
	
	addOneByPointer(&count)
	fmt.Printf("Count:\t %d | Addr: %v\n", count, &count)
}

func addOneByValue(inc int) {
	inc++
	fmt.Printf("AddOne:  %d | Addr: %v\n", inc, &inc)
}

func addOneByPointer(inc *int) {
	*inc++
	fmt.Printf("*AddOne: %d | Addr: %v\n", *inc, inc)
}
---
Output:
Count:	 10 | Addr: 0x1a0008

AddOne:  11 | Addr: 0x414040
Count:	 10 | Addr: 0x1a0008

*AddOne: 11 | Addr: 0x1a0008
Count:	 11 | Addr: 0x1a0008
```

[Go Playground Example...](https://play.golang.org/p/6VmfG8U7v2-)

In the above example, we can see that we have 3 different program boundaries 
(main, addOneByValue, and addOneByPointer). Each of these functions can access data that is 
stored in memory local to its boundary. 

The count variable is stored in a memory location within main's program boundary (`0x1a0008`). We can 
manipulate the data in this location anyway we want within the boundaries. 

When we pass the variable `count` by-value into the `addOneByValue` function, we are taking the value of count, 
and passing it into the next program boundary. We are effectively saying, "take the value of `10`, and assign
it to a memory location that is local to your program boundary". The next function will take this value, and
create a copy of it. 

When we pass the variable `count` by-pointer into the `addOneByPointer` function, we are passing the memory 
address of `0x1a0008` in and saying "treat `0x1a0008` as a valid memory address, and use this address location
to find the value of `count`.  By doing this, we are able to share the variable `count` across program boundaries.

#### Literal Values and Pointers
An interesting note about using values as pointers. Literal values are considered constants in Go and
only exist and evaluated at compile time. This means that there is no address value for a literal value 
and the data cannot be shared. If we were to run the following code 
```go
import "fmt"

func main() {
	test := &10
	fmt.Println(test)
}
```
We would get an error stating `prog.go:6:10: cannot take the address of 10`. This is because `10` is 
a constant in Go.

---


## Garbage Collection in Go
As pointers and other declarations happen in code, memory gets allocated and 
the GC's job is to determine which of these memory allocations cannot or will not
be used anymore, and can be deallocated.

In other words, it's job is to free up memory that is no longer needed. I don't want to go
into too much detail about the GC here, but if you are interested, please have a look at these docs:
+ [The Journey of Go's Garbage Collector](https://blog.golang.org/ismmkeynote)
+ [Garbage Collection in Go](https://www.ardanlabs.com/blog/2018/12/garbage-collection-in-go-part1-semantics.html)

Some key points about the GC:
 1. Garbage collection happens automatically, behind the scenes.
 2. GC runs when the memory allocation since last GC has reached the `GOGC` variable (set to 100% by default).
 Using default values, if the memory allocated is 5MB, then GC will kick off again when the memory allocation has reached 100% 
 more allocation, or 10MB.
 3. To debug GC activity, you can run your application with an env var of `GODEBUG=gctrace=1`
 4. Keeping heap memory allocation small will cause the GC to run less often.
    * Think reference types: slices, maps, interfaces, functions, channels, and type pointers. 