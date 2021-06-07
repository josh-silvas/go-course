// Modules Example: In this example, we use the same packaging and imports from the package
//    example, however, we import additional 3rd party packages we want to manage via modules.
//
//
// Before you begin:
//   1. Change to the practice directory
//       - cd practice/packaging/02-practice-modules
//   2. Execute `go run main.go`

package main

import (
	"fmt"

	"github.com/kyokomi/emoji"
	"github.com/josh5276/go-course/practice/packaging/01-example-pkg/fibonacci"
)

const (
	startSeq = 40
	endSeq   = 45
)

// TODO: 1. Make sure Go Modules is enabled by setting `export GO111MODULE=auto`
// TODO: 2. Import the new package using `go get github.com/kyokomi/emoji`
// TODO: 3. Run `go mod init` to initialize this package
// TODO: 4. View the contents of go.mod `cat go.mod`
// TODO: 5. Run the application using `go run main.go`

func main() {
	fmt.Println("Running fibonacci sequence using recursion")
	for i := startSeq; i <= endSeq; i++ {
		fmt.Println(emoji.Sprintf(":stop_sign: fib#%d = %d", i, fibonacci.CalcSlow(i)))
	}

	fmt.Println("Running fibonacci sequence using iteration")
	for i := startSeq; i <= endSeq; i++ {
		fmt.Println(emoji.Sprintf(":high-speed_train: fib#%d = %d", i, fibonacci.CalcFaster(i)))
	}
}
