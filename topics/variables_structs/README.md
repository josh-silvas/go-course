## Variables and Structs
Variables and structs are the foundation of Go's type system. They are used in
every interaction between the code you are writing (data manipulation) and the 
memory location of that data. 

Think of variables as the DNS system of the language. Where `www.google.com` is a
reference to `142.250.114.104`, `myVarName` can be a short-hand reference to `0x1a0008`

Structs and variables are defined as a certain type and the compiler expects those
types to be accessed as their definition, making Go's access to memory [type-safe](https://github.com/josh5276/go-course/tree/master/reference#type-safety).

**Variable Initialization**

There are two primary ways to initialize a variable in Go, either one is acceptable methods.
_Use consistent initialization methods_

Always use the `var` keyword when declaring variables at their empty values
```$xslt
var t []string
```

When declaring and assigning, you can use either method, as both are valid.
```$xslt
var t = []string{"one", "two"}
t := []string{"one", "two"}
```

**Variable Names**

Effective Go prefers short variable names over long ones. This is especially 
true for local variables. Prefer `c` over `lineCount`. Prefer `i` over 
`index`. Common variables such as loop indices and readers 
can be a single letter, such as `i` or `r`. 

_Ref: CodeReviewComments#Variables_

#### References
+ [Effective Go#Variables](https://golang.org/doc/effective_go.html#variables)
+ [CodeReviewComments#Variables](https://github.com/golang/go/wiki/CodeReviewComments#variable-names)

---


## Constants and Conversions
Constants are memory locations that hold a set of data that is evaluated at compile time and never change.
+ If an untyped constant is declared, you can assign it to a type later, however, you cannot reassign a typed constant without an explicit conversion.
+ Constants will perform implicit type conversion at compile time.
+ In the example below, the compiler will implicitly type `5` as an integer, however, it can be typed as a another numerical type that is within the type boundaries.
```$xslt
const (
	unTyped = 5
	typed float64 = 5
)
func main() {
	fmt.Printf("Type %T: %v\n", unTyped, unTyped)
	fmt.Printf("Type %T: %v\n", typed, typed)
}
----
Type int: 5
Type float64: 5
```
+ Explicit type conversions are ones that we are coding to purposely perform a type-conversion, outside of the compiler. For example, the following snippet will result in an error because the time method `Add()` requires a type of `time.Duration`. 
```$xslt
const typed float64 = 5

func main() {
	fmt.Printf("Type %T: %v\n", typed, time.Now().Add(typed))
}
----
cannot use typed (type float64) as type time.Duration in argument to time.Now().Add
```
An example of converting this type explicitly would be to set the type of `time.Duration` on the constant. 
Likewise, we could have implicitly allowed Go to convert the type by not specifying a type at declaration.
```$xslt
const typed float64 = 5

func main() {
	fmt.Printf("Type %T: %v\n", typed, time.Now().Add(time.Duration(typed)))
}
```
[Go Playground](https://play.golang.org/p/TXKhSkK0u72)

---

## Structs and Complex Types
A struct is a complex literal definition of smaller groups of data, in which we want to organize in a way
that the compiler and system allocate memory for in an alignment boundary to minimize memory defragmentation.

#### Defining Struct Elements
A struct can be defined using a named field called an IdentifierList, 
or implicitly included (called an EmbeddedField).

Identifier elements are pretty straightforward. You can define an identifier element by naming the type you are 
creating within the struct. In the example below, we are creating two elements called `First` and `Last`
```$xslt
type Name struct {
    First string
    Last string 
}
``` 

Type embedding allows a type to import another type to give it access to any of its types and methods that are 
exported. Go calls this type-promoting. Using embedding, you can take the above example of `type Name struct {` and 
embed it into another structure to inherit it's properties. For example...
```$xslt
type Person struct {
	Name
	Age int
}


func main() {
	var person    = Person{Age: 37}
	person.First = "This"
	person.Last = "OldDude"
	
	fmt.Printf("%s %s is %d old.", person.First, person.Last, person.Age)
}
---
Output:
This OldDude is 37 old.
```
We are able to access the named element of `First` and `Last` within the `Person` struct by embedding the `Name` 
struct within it.


[Go Playground Example](https://play.golang.org/p/RV0DrweFSWC)

#### Struct Memory Allocation
Using a struct of in the example below, we see that the declared types have a memory size of 
```$xslt
int16:    02 bytes
int:      04 bytes
float64:  08 bytes
[]string: 12 bytes
```
Which would add up to a total of 26 bytes, however, since memory is allocated to types in a way so that the language
minimizes memory defraging, we can see that the below structure actually gets padded where needed to fit the systems 
alignment boundaries.
```
type Example struct {
	int16Var int16
	intVar   int
	floatVar float64
	sliceVar []string
}
---
Output:
System alignment boundary: 4
Struct is 32 bytes long
int16Var | Offset:00 | Size:02
  intVar | Offset:04 | Size:04
floatVar | Offset:08 | Size:08
sliceVar | Offset:16 | Size:12
```
Because the alignment boundary is in increments of 4 bytes, and the first type (int16) only allocates 2 bytes, the memory allocation
will pad 2 bytes between `int16Var` and `intVar`. This means that we could technically allocate another `int16` type between
`int16Var` and `intVar` and the system will still only allocate 32 bytes of memory for the structure.
```$xslt
type Example struct {
	int16Var int16
	int162Var int16
	intVar   int
	floatVar float64
	sliceVar []string
}
---
Output:
System alignment boundary: 4
Struct is 32 bytes long
int16Var  | Offset:00 | Size:02
int162Var | Offset:02 | Size:02
  intVar  | Offset:04 | Size:04
floatVar  | Offset:08 | Size:08
sliceVar  | Offset:16 | Size:12
```

[Go Playground Example...](https://play.golang.org/p/6PqRHA1_cpp)
#### References
 [Type Embedding in Go](https://travix.io/type-embedding-in-go-ba40dd4264df)
 