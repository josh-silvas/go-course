## Arrays/Slices
Slices in Go are an abstracted layer on top of Go's original array design. 

#### Arrays
Arrays are fixed-length sequences of items of the same type. arrays in Go are values. 
This has a couple of important implications: 
  1. Assigning one array to another copies all of the elements
  2. If you pass an array to a function, it will receive a copy of the array (not a pointer or reference to it).

The syntax for creating an array is similar to a slice, however, you must supply an array size at initialization.
``go
var myArray := [5]int{1, 2, 3, 4, 5}
``

Alternatively, you can allow the compiler to count the array size needed, although, this is still not dynamic.
``go
var myArray := [...]int{1, 2, 3, 4, 5}
``

Both of these initialization methods will create a 5 element array with a fixed size and attempting to access
an index outside of this decalred size will result in an error.
```go
fmt.Println(myArray[5]) 
// invalid array index 5 (out of bounds for 5-element array)
```

#### Slices
Slices offer a much more flexible wrapper on top of arrays and are generally the accepted implementation throughout
Go code.

Some highlights of slices functionality:
 1. Slices can be resized using the built-in `append()` function. 
 2. Slices are reference types (think pointers), meaning that they are cheap to assign and can be passed to other 
 functions without having to create a new copy of its underlying array. 
 3. The functions in Go’s standard library all use slices in favor of arrays.
 4. **When slices exceed their initial memory allocation, the append function doubles 
 capacity each reallocation for the first few thousand elements, and then 
 progresses at a rate of ~1.25 capacity growth after that.
 
_** William Kennedy verified behavior [here](https://www.ardanlabs.com/blog/2013/08/collections-of-unknown-length-in-go.html)._
 
##### Slice Declaration
Slice declaration can be done using the literal assignment:
```go
var mySlice = []int{1, 2, 3}
```
Or by using the [make](https://github.com/josh5276/go-course/tree/master/reference#the-make-built-in) built-in function.
```go
var mySlice = make([]int, 0)
```

##### Slice of a Slice
Using the `:` operator, we are able to select start and/or end positions of a slice to retrieve only 
a portion of the data. The syntax is `[start_position:end_position]`, where an excluded value on either
side represents the max value for that position. For example:
```go
func main() {
	mySlice := []string{"a", "b", "c", "d", "e", "f", "g"}

	// Print from the beginning position to the end
	fmt.Println(mySlice[:])

	// Print from position 2 to the end position
	fmt.Println(mySlice[2:])

	// Print from the beginning position to position 3
	fmt.Println(mySlice[:3])

	// Print items between position 1 to position 3
	fmt.Println(mySlice[1:3])
}

Result:
-----------------
[a b c d e f g]
[c d e f g]
[a b c]
[b c]
```
+ [Go Playground Example](https://play.golang.org/p/MpOOEgW6Gmb)

##### Merging Slices
Go provides an easy way to take two slices and merge them together into a single slice.
The same underlying logic applies to merging slices, regarding copying to a new slice
to grow the size, if needed.

To merge a slice, use the `...` syntax within the `append` function:
```go
mySlice1 := []int{1, 2, 3}
mySlice2 := []int{4, 5, 6}
mySlice1 = append(mySlice1, mySlice2...)
fmt.Println(mySlice1)

Result:
-----------------
[1 2 3 4 5 6]
```
+ [Go Playground Example](https://play.golang.org/p/y3HHJA6YeYi)

##### Slice Implementations
It's worth noting that slices are used in some of the fundamental parts of the language. Strings, for example,
are nothing more than a slice of bytes that are immutable (cannot be changed). So if we were to look at the 
composure of a string and a slice, it might look something like:
```go
String:         Slice:
+------+        +------+
| data |        | data |
+------+        +------+
| len  |        | len  |
+------+        +------+
                | cap  |
                +------+
```
Since strings are immutable, the len is always the cap. 

Now knowing that a string is almost equivalent to a slice of bytes, and we know that with slices we are able to 
take a slice of a slice, we can perform similar actions with strings.

For example:
```go
fmt.Println("ORD1"[:3])

Result:
-----------
ORD
```


#### References
+ [Arrays vs Slices](https://www.godesignpatterns.com/2014/05/arrays-vs-slices.html)
+ [Go Slices Usage and Internals](https://blog.golang.org/go-slices-usage-and-internals)
