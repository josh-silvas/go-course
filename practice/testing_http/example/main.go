// Unit Testing Answer: In this example, we importing a function from the package and
// implementing a test against it.
// Before you begin:
//   1. Change to the practice directory
//       - cd practice/testing_unit
//   2. Complete the `TO DO` sections of the code
//   3. Check answers in answer/main_test.go

package testing_http

import (
	"encoding/json"
	"net/http"
	"strings"
)

type User struct {
	FName string `json:"first_name"`
	LName string `json:"last_name"`
	Email string `json:"email"`
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

// Route contains all the http handlers for this API
func Route() {
	http.HandleFunc("/user/", GetUser)
}

// GetUser represents the HTTP version of the GetUser function from the
// unit test example.
func GetUser(rw http.ResponseWriter, r *http.Request) {
	var resp User

	// This is terrible. There are so many http packages that will handle this better.
	// This test only uses the net/http package to show the low-level interactions without
	// having to import a 3rd part package.
	user := strings.TrimPrefix(r.URL.Path, "/user/")

	for _, u := range users {
		if strings.EqualFold(u.FName, user) {
			resp = u
			break
		}
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(rw).Encode(&resp); err != nil {
		panic(err)
	}
}
