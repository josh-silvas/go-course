// Example usage of the standard Go error implementation.
// We can use custom error variables to test error types returned
package main

import (
	"errors"
	"strconv"

	"github.com/sirupsen/logrus"
)

var (
	// Defining the different type of error messages that can be returned by
	// the isEven function.
	// NOTE: Because of Go's type system, the error message here does not matter in
	// differentiating between the errors received. Go will create a new instance
	// of type error for each declaration, and the comparison is done against the
	// instance.
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
