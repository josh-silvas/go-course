// Functions Example: In this example, we will be creating a new function that timestamps a
// log message to terminal. We will call our function at the beginning and end of main()
// Before you begin:
//   1. Change to the practice directory
//       - cd practice/functions
//   2. Complete the `TO DO` sections of the code
//   3. Check answers in answer/main.go

package main

import "time"

// TODO: 1. Complete the requirements for the logger function
// The logger function should take a "message" string, then call the log.Printf function to
// print the passed in "message".
//
// Note: The log.Printf function will prepend on a timestamp.
func logger() {
}

func main() {
	// TODO: 2. Call the logger function to print the current timestamp

	// TODO: 3. Call the same logger function using a defer statement

	// time.Sleep will force the program to wait for a given duration
	// before continuing.
	time.Sleep(5 * time.Second)
}
