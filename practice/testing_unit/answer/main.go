// Unit Testing Answer: In this example, we importing a function from the package and
// implementing a test against it.
// Before you begin:
//   1. Change to the practice directory
//       - cd practice/testing_unit
//   2. Complete the `TO DO` sections of the code
//   3. Check answers in answer/main_test.go

package unit_test_answer

import "strings"

type User struct {
	FName string
	LName string
	Email string
}

var users = []User{
	{
		FName: "Shooter",
		LName: "McGavin",
		Email: "shooter.mcgavin@example.com",
	},
	{
		FName: "Go",
		LName: "Pher",
		Email: "gopher@golang.org",
	},
}

// GetUser is an example function exported so that the testing package
// can consume this function.
func GetUser(user string) (res User) {
	for _, u := range users {
		if strings.EqualFold(u.FName, user) {
			res = u
			break
		}
	}
	return
}

// GetEmail is an example function exported so that the testing package
// can consume this function.
func GetEmail(email string) (res User) {
	for _, u := range users {
		if strings.EqualFold(u.Email, email) {
			res = u
			break
		}
	}
	return
}
