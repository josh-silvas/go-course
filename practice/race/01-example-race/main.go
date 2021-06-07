// Example file containing a function that has a designed race in the body.
//   go run -race main.go
//
// To view this file as assembly, build the executable and use Go compile tool to output to file.
//   go tool compile -S main.go > main.s

package main

import (
	"fmt"
	"sync"
)

var (
	// countVar is the global variable we will use to trigger a race condition.
	// Access to this memory location from multiple routines will trigger a race.
	countVar = 0

	// We use a wait group here to allow the routines to finish before we
	// proceed on to printing the results.
	wg sync.WaitGroup
)

func main() {

	// Loop through and create goroutines to launch the "count()" function.
	// This iteration will start two separate routines that will ultimately
	// access the same memory location unsafely.
	for r := 1; r <= 2; r++ {
		wg.Add(1)
		go count()
	}

	// Using a wait group, we can wait here until all programs have
	// finished executing.
	wg.Wait()

	fmt.Printf("total value of %d\n", countVar)
}

// count examples a function that both reads and writes to the same variable. Using
// this function in concurrent routines without protecting access to countVar, can
// result in a race condition.
func count() {
	for count := 0; count < 4; count++ {

		// Read the countVar
		tempVar := countVar

		// Increment the temporary variable
		tempVar += 1

		// Write the new value back to the countVar
		countVar = tempVar

	}
	wg.Done()
}
