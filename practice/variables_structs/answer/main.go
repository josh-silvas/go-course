// Variables and Structs Answer: In this example, we will create a set of variables at their empty
//    values, then assign values to them, and print them.
// Before you begin:
//   1. Change to the practice directory
//       - cd practice/variables_structs
//   2. Complete the `TO DO` sections of the code
//   3. Forget that the month born is not taken into consideration
//   4. Check answers in answer/main.go

package main

import "fmt"

const (
	// TODO: 1. Define a constant named 'age' and assign a value to it
	year = 2019
)

var (
	// TODO: 2. Define 2 variables named `first` and `last` of type string and set to their empty values.
	first string
	last  string

	// TODO: 3. Define 1 variables named `age` of type int and set to its empty value.
	age int
)

func main() {
	// TODO: 4. Assign values to `first` and `last` strings, and the `age` int.
	first = "Josh"
	last = "Silvas"
	age = 37

	// TODO: 5. Subtract `age` variable from the `year` constant, and assign to a yearBorn value.
	yearBorn := year - age

	// TODO: 6. Print the first, last variables and the yearBorn.
	fmt.Printf("%s %s was born in %d\n", first, last, yearBorn)
}

// Results:
// ---------------------------------------------------------------------
// Josh Silvas was born in 1982
