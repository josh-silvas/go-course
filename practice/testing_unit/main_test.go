// Unit Testing Example: In this example, we importing a function from the package and
// implementing a test against it.
// Before you begin:
//   1. Change to the practice directory
//       - cd practice/testing_unit
//   2. Complete the `TO DO` sections of the code
//   3. Test the program using the ` go test -v -bench=.` syntax
//   4. Check answers in answer/main_test.go

package unit_test_example

import (
	"testing"
)

func TestGetUser(t *testing.T) {
	// TODO: 1. Call the GetUser function using the name "shooter". Store the results in a variable.

	// TODO: 2. Test if the last name (LName) is equal to "McGavin". If not, fail the test.

	// TODO: 3. Log a message stating that the test completed successfully
}

func ExampleGetEmail() {
	// TODO: 4. Create a similar test using the example method on the GetEmail function
}

// TODO: 5. Define a package level variable to assign benchmark results to

func BenchmarkGetUser(b *testing.B) {
	var r User
	for n := 0; n < b.N; n++ {
		// TODO: 6. Run the GetUser function and store them in a variable local to the program boundary.
	}

	// TODO: 7. Store the final result to the package level variable.
}
