## Error Handling
##### Creating a Custom Error
To create a customer error type, we can create a new type that implements the error interface. To match the error interface
signature, we need to create a type that implements a method called `Error()` that returns a `string`. The example below 
demonstrates a custom error type called NotEvenError (idiomatic Go states that customer errors should end with `Error`).


```go
type NotEvenError struct {
	// Path at which the error occurred
	Path string

	// Err is the error that occurred during the operation.
	Err error
}

func (e *NotEvenError) Error() string {
	if e == nil {
		return "<nil>"
	}
	return fmt.Sprintf("%s:%s", e.Path, e.Err)
}
```

We can now use this new custom error as a more granular error type in our package. We can test against the errors
concrete type to determine the error we received. 

Below is an example of looping through strings and executing the `isEven` function on them. The returning error is then
tested against its type to determine if the error received is of type `NotEvenError`

```go
func main() {
	for _, s := range []string{"52", "15", "asd"} {
		if err := isEven(s); err != nil {
			switch err.(type) {
			case *NotEvenError:
				logrus.Warnf("i just can't even:%s", err)
			default:
				logrus.Errorf("unknown error:%s", err)
			}
			continue
		}
		logrus.Infof("%s is a even number", s)
	}
}

func isEven(numberStr string) error {
	numberInt, err := strconv.Atoi(numberStr)
	if err != nil {
		return err
	}
	if numberInt%2 != 0 {
		return &NotEvenError{Path: "main:isEvenCheck", Err: errors.New("odd number")}
	}
	return nil
}
```

[Source Example](https://github.com/josh5276/go-course/blob/default/topics/error_custom/example.go)

[Go Playground Example](https://play.golang.org/p/P5XFLO8CDEX)