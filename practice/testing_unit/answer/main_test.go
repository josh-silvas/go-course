// Unit Testing Answer: In this example, we importing a function from the package and
// implementing a test against it.
// Before you begin:
//   1. Change to the practice directory
//       - cd practice/testing_unit
//   2. Complete the `TO DO` sections of the code
//   3. Check answers in answer/main_test.go

package unit_test_answer

import (
	"fmt"
	"strings"
	"testing"
)

func TestGetUser(t *testing.T) {
	user := GetUser("shooter")
	if !strings.EqualFold(user.LName, "mcgavin") {
		t.Fatalf("last name of %s unexpected", user.LName)
	}
	t.Logf("successfully found %s %s", user.FName, user.LName)

	user = GetUser("go")
	if !strings.EqualFold(user.LName, "pher") {
		t.Fatalf("last name of %s unexpected", user.LName)
	}
	t.Logf("successfully found %s %s", user.FName, user.LName)
}

// http://localhost:8080/pkg/github.com/josh5276/go-course/practice/testing_unit/answer/
func ExampleGetEmail() {
	user := GetEmail("shooter.mcgavin@example.com")
	fmt.Println(user.FName)
	fmt.Println(user.LName)
	// Output:
	// Shooter
	// McGavin
}

var benchUserResult User

func BenchmarkGetUser(b *testing.B) {
	var r User
	for n := 0; n < b.N; n++ {
		r = GetUser("shooter")
	}
	benchUserResult = r
}

// Results:
// ---------------------------------------------------------------------
// === RUN   TestGetUser
// --- PASS: TestGetUser (0.00s)
// main_test.go:24: successfully found Shooter McGavin
// main_test.go:30: successfully found Go Pher
// === RUN   ExampleGetEmail
// --- PASS: ExampleGetEmail (0.00s)
// goos: darwin
// goarch: amd64
// BenchmarkGetUser-4   	47299801	        25.0 ns/op
// PASS

// go test -v -coverprofile cover.out
// go tool cover -html=cover.out
