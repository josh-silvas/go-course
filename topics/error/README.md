## Error Handling
Error handling in Go is implemented using the `errors` package built into the language. Each function can 
return a type of _error_ to allow the calling program to evaluate if an error occurred during processing. 

Some idiomatic points of Go's error handling:
 * Typically, you can return either an error, or an error and response as your return values.
 This allows for the calling function to evaluate if an error was received.
    
    `func Atoi(s string) (int, error) {` 
    
 * When possible, create package level error values. This allows for calling functions to compare error responses
 to make intelligent decisions about the error received.
 
    `var ErrInvalidVar = errors.New("mypkg: invalid variable passed in")`
 
 * If the standard error package does not provide enough context to make an intelligent decision, create custom 
 error types by implementing the error interface.
 
 * Error values should be named with a prepending `Err`, such as `ErrInvalidVar`
 * Custom error implementations should be named with an appending `Error`, such as `CustomError`

##### Using the Standard Error Package
Golang standard error package:
```go
// Error interface defined in built-ins
//
// The error built-in interface type is the conventional interface for
// representing an error condition, with the nil value representing no error.
type error interface {
	Error() string
}

// ----------------------------------------------------------------------
// Errors package
package errors

// New returns an error that formats as the given text.
func New(text string) error {
	return &errorString{text}
}

// errorString is a trivial implementation of error.
type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}
```

The standard error package that's built into the Go ecosystem is suitable for most use-cases. We can see that it creates
an `errorString` type, and implements the method Error() to return a string value of the type. This satisfies the 
error interface.

Here is an example of implementing Go's standard error package and using a pre-defined set of error variables to 
return for the isEven function. 

The neat part about defining your errors as variables, is that the creation of the error using `errors.New()` will create
a new error pointer for each instantiation. Meaning that the reference location of the errors are compared and not the
text in the error message, meaning error messages can be changed without adverse effects in error testing.

```go
var (
	ErrNotEven   = errors.New("main: odd number passed in")
	ErrNotNumber = errors.New("main: non-integer passed in")
)

func main() {
	for _, s := range []string{"52", "15", "asd"} {
		if err := isEven(s); err != nil {
			switch err {
			case ErrNotNumber:
				logrus.Errorf("%s is not a number bruh", s)
			case ErrNotEven:
				logrus.Errorf("%s is an odd number", s)
			}
			continue
		}
		logrus.Infof("%s is a even number", s)
	}
}

func isEven(numberStr string) error {
	numberInt, err := strconv.Atoi(numberStr)
	if err != nil {
		return ErrNotNumber
	}
	if numberInt%2 != 0 {
		return ErrNotEven
	}
	return nil
}
```

[Source Example](https://github.com/josh5276/go-course/blob/default/topics/error/example.go)

[Go Playground Example](https://play.golang.org/p/LN-wTP2Rbnp)