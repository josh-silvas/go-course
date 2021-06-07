// Concurrency Practice: In this example, we are going to call out to an api using Goroutines.
// Before you begin:
//   1. Change to the practice directory
//       - cd practice/concurrency
//   2. Complete the `TO DO` sections of the code
//   3. Check answers in answer/main.go

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"
)

func main() {
	// TODO: 1. Create an empty variable of type sync.WaitGroup

	// TODO: 2. Add 2 deltas to the wait group for the two Goroutines

	// TODO: 3. Launch a Goroutines calling the 'callPolygon()' function, passing it a pointer to the wait group

	// TODO: 4. Launch a Goroutines calling the 'callTechCrunch()' function, passing it a pointer to the wait group

	// TODO: 5. Wait for the wait group to fully decrement before proceeding
}

func callPolygon(wg *sync.WaitGroup) {
	call("polygon", wg)
}

func callTechCrunch(wg *sync.WaitGroup) {
	call("techcrunch", wg)
}

// call function actually handles the API call and response unmashalling.
func call(source string, wg *sync.WaitGroup) {
	defer wg.Done()
	var data struct {
		Status       string `json:"status"`
		TotalResults int    `json:"totalResults"`
	}
	req, err := http.NewRequest(http.MethodGet, "https://newsapi.org/v2/top-headlines", nil)
	if err != nil {
		return
	}
	req.URL.RawQuery = url.Values{
		"apiKey":  []string{"ed68824cbc024ea796df56de491c3cb0"},
		"sources": []string{source},
	}.Encode()

	doResp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}

	// The response from this API returns quickly. Lets simulate an API that may take longer to respond
	time.Sleep(3 * time.Second)
	if err := json.NewDecoder(doResp.Body).Decode(&data); err != nil {
		return
	}
	fmt.Printf("%s found %d articles from %s\n", time.Now().Format(time.RFC1123), data.TotalResults, source)
}

// Results:
// ---------------------------------------------------------------------
// Sun, 08 Dec 2019 20:36:13 CST found 10 articles from techcrunch
// Sun, 08 Dec 2019 20:36:13 CST found 10 articles from polygon
