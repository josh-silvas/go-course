// Example usage of customer error messaging interface implementations.
package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/sirupsen/logrus"
)

type NotEvenError struct {
	// Path at which the error occurred
	Path string
	Err  error
}

func (e *NotEvenError) Error() string {
	if e == nil {
		return "<nil>"
	}
	return fmt.Sprintf("%s:%s", e.Path, e.Err)
}

type InvalidFormatError struct {
	// Path at which the error occurred
	Path string
	Err  error
}

func (e *InvalidFormatError) Error() string {
	if e == nil {
		return "<nil>"
	}
	return fmt.Sprintf("%s:%s", e.Path, e.Err)
}

func main() {
	numberStr := []string{"52", "15", "asd"}

	// Range through the list of numbers and run the isEven check
	// on them. Digest any error types that are returned.
	for _, s := range numberStr {
		if err := isEven(s); err != nil {
			// Switch case against the type of error that was received. Here we can
			// test against the custom concrete types we created for errors.
			switch err.(type) {
			case *NotEvenError:
				logrus.Warnf("i just can't even:%s", err)
			case *InvalidFormatError:
				logrus.Error(err)
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
		// If we run into an error converting the string to an integer,
		// return a custom type of InvalidFormatError
		return &InvalidFormatError{Path: "main", Err: err}
	}
	if numberInt%2 != 0 {
		// If the number is odd, return a custom error type of NotEvenError
		return &NotEvenError{Path: "main:isEvenCheck", Err: errors.New("odd number")}
	}
	return nil
}
