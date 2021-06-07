// Functions Answer: In this example, we will be creating a new function that timestamps a
// log message to terminal. We will call our function at the beginning and end of main()
// Before you begin:
//   1. Change to the practice directory
//       - cd practice/functions
//   2. Complete the `TO DO` sections of the code
//   3. Check answers in answer/main.go

package main

import (
	"log"
	"time"
)

// The logger function should take a "message" string, then call the log.Printf function to
// print the passed in "message".
//
// Note: The log.Printf function will prepend on a timestamp.
func Logger(message string) {
	log.Printf("type=INFO, message=%s", message)
}

func main() {
	Logger("Beginning of execution.")

	defer Logger("from deferred function")

	// time.Sleep will force the program to wait for a given duration
	// before continuing.
	time.Sleep(5 * time.Second)
}

// Results:
// ---------------------------------------------------------------------
// 2019/11/12 22:03:34 type=INFO, message=begging of execution
// 2019/11/12 22:03:39 type=INFO, message=from deferred function
