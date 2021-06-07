// Interfaces Answer: In this example, we will be defining an interface that implements an
//   age method. The interface will convert age into specific typed ages.
//
// Before you begin:
//   1. Change to the practice directory
//       - cd practice/interfaces
//   2. Complete the `TO DO` sections of the code
//   3. Check answers in answer/main.go

package main

import (
	"fmt"
)

// TODO: 1. Define an interface 'Ager' with one method called 'age'
type Ager interface {
	age()
}

// TODO: 2. Define a type named Human of type int
type Human int

// TODO: 3. Define a method called 'age()' that takes Human as a value
// TODO:    receiver. Have the method print the Age value.
func (h Human) age() {
	fmt.Printf("I am %d years old in human years.\n", h)
}

// TODO: 4. Define a type named Dog of type int
type Dog int

// TODO: 5. Define a method called 'age()' that takes Dog as a value
// TODO:    receiver. Have the method print the age * 7 value.
func (d Dog) age() {
	fmt.Printf("I am %d years old in dog years.\n", d*7)
}

// printAge function takes the Ager interface and abstracts the
// ager functionality.
// TODO: 6. Call the age method inside of the printAge function.
func printAge(a Ager) {
	a.age()
}

func main() {
	// TODO: 7. Initialize a Human type with an age and pass it into the printAge function.
	printAge(Human(37))

	// TODO: 8. Initialize a Dog type with an age and pass it into the printAge function.
	printAge(Dog(37))

	// Alternatively, we can initialize the type and call the method directly.
	// Human(37).age()
	// Dog(10).age()
}

// Results:
// ---------------------------------------------------------------------
// I am 37 years old in human years.
// I am 259 years old in dog years.
