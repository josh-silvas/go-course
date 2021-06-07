## Functions
Functions are another core piece of the Go language that allows you to group code together in
a reusable manner. Some high points regarding functions:

 - Functions starting with an uppercase letter can be exported and used in another package. Function names
starting with a lowercase letter are considered to be internal to the package and cannot be exported.
 - Functions can (and usually will) return multiple values. It's common practice in the Go language to 
 return a value and error. Ex. `func myFunc(value string) (string, error)`
 - Naming return values will create a local variable to the function, and do not need to be redeclared. For example:
    ```go
    func myFunc() (num int, err error) {
       num = 10
       return
    }
    ```
 - Functions can be anonymous when there is a need to group code together for a single use. Defer 
 statements are a good example of a use-case for anonymous functions.

##### Deferred Functions
Go has a builtin statement (`defer`) to handle scenarios where a function call should 
run immediately before the function executing the statement returns. 

This offers an effective way to execute a function when a parent function ends,
regardless of the state of the program that happens within the parent function.

NOTE: arguments, or receivers that are passed in or used by the deferred function
are evaluated at deferred execution time, not the call execution. For example:
```go
desc := "anonymous func call"

defer func(name string) {
	fmt.Printf("printing in the %s statement\n", name)
}(desc)

func(name string) {
	fmt.Printf("printing in the %s statement\n", name)
}(desc)

desc = "deferred statement"

Result:
------------------
printing in the anonymous func call statement
printing in the anonymous func call statement
```
In the above example, we define the description variable as `anonymous func call`, then 
we defer a function and call a function directly, passing in the `desc` variable. Finally,
we modify the `desc` variable to `deffered statement`. Although the deferred function
is not actually executed until the program exits, the variable passed into the function
is evaluated when the defer statement is executed.

##### Frame Boundaries
Functions execute within the scope of frame boundaries that provide an individual
memory space for each respective function. Each frame allows a function to operate 
within their own context and also provides flow control. 

A function has direct access to the memory inside its frame, through the frame 
pointer, but access to memory outside its frame requires indirect access. For a 
function to access memory outside of its frame, that memory must be shared with 
the function. The mechanics and restrictions established by these frame boundaries 
need to be understood and learned first.

#### References
+ [Effective Go#Functions](https://golang.org/doc/effective_go.html#functions)
+ [Pointer Mechanics](https://www.ardanlabs.com/blog/2017/05/language-mechanics-on-stacks-and-pointers.html)
