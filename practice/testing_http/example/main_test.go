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
	"net/http/httptest"
	"testing"
)

// init function runs prior to main in any execution. This function
// will serve the routes for testing purposes.
func init() {
	Route()
}

func TestRoute(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/user/shooter", nil)
	w := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Fatalf("unexpected response code - %d", w.Code)
	}

	var resp User
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("unable to decode response from API. Err:%s", err)
	}
	if resp.LName == "McGavin" {
		t.Logf("successfully found %s %s", resp.FName, resp.LName)
	} else {
		t.Errorf("unable to find user")
	}
}

// Results:
// ---------------------------------------------------------------------
// === RUN   TestRoute
// --- PASS: TestRoute (0.00s)
// main_test.go:38: successfully found Shooter McGavin
// PASS
