// Packaging Example: In this example, we use a custom package we created to calculate the fibonacci
//    sequence.
//
// Before you begin:
//   1. Change to the practice directory
//       - cd practice/packaging/01-example-pkg
//   2. Execute `go run main.go`

package main

import (
	"fmt"

	"github.com/josh5276/go-course/practice/packaging/01-example-pkg/fibonacci"
)

const (
	startSeq = 35
	endSeq   = 40
)

func main() {
	for i := startSeq; i <= endSeq; i++ {

		fmt.Printf("fib#%d = %d", i, fibonacci.CalcSlow(i))
	}
	for i := startSeq; i <= endSeq; i++ {
		fmt.Printf("fib#%d = %d", i, fibonacci.CalcFaster(i))
	}
}
