# Golang Training References

#### Type-Safety
Type-Safety can be defined by a programming language that prevents actions to
be performed on a set of data that is not allowed by the data's type. Go achieves 
type-safety by requiring hard-typed data to be defined before actioning on said 
data. For instance, if a integer is typed as a string `var i = "5"`, you cannot 
access methods that expect a integer in it's computation

[Go Playground Example...](https://play.golang.org/p/XehOciwvkRi)

#### The Make Built-in
The make built-in function allocates and initializes an object of type
slice, map, or chan (only). The return type is the same as the type of its
argument, not a pointer to it. The specification of the result depends on
the type:
 * Slice: The size specifies the length. The capacity of the slice is
   equal to its length. A second integer argument may be provided to
   specify a different capacity; it must be no smaller than the
   length. For example, `make([]int, 0, 10)` allocates an underlying array
   of size 10 and returns a slice of length 0 and capacity 10 that is
   backed by this underlying array.
 * Map: An empty map is allocated with enough space to hold the
   specified number of elements. The size may be omitted, in which case
   a small starting size is allocated.
 * Channel: The channel's buffer is initialized with the specified
   buffer capacity. If zero, or the size is omitted, the channel is
   unbuffered.

[Ref: Go Builtin Source Code](https://golang.org/src/builtin/builtin.go)
   
#### The New Built-in
The new built-in function allocates memory. The first argument is a type,
not a value, and the value returned is a pointer to a newly
allocated zero value of that type.

[Ref: Go Builtin Source Code](https://golang.org/src/builtin/builtin.go)