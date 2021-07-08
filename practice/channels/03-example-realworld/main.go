package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/mkideal/pkg/errors"
)

// Current Rate limit:
const rateLimit = time.Millisecond * 600

func main() {
	// Parse list of devices from a CSV file
	devices := parseCSV()

	// NewTicker will create a Ticker type with two important things in it.
	//  1. A channel of type time.Time with a 1 element buffer 'make(chan time.Time, 1)'
	//  2. A runtime Timer that holds the duration, next iteration, and increment func
	//
	// When time reaches 'duration', the timer will send the current time (time.Now())
	// on it's channel. This allows the receiving function to receive the current time,
	// unblocking it for an iteration.
	//
	// A 1 element buffered channel guarantees that the previous signal that was sent has been received.
	// Because the receive of the first signal, happens before the send of the second signal completes.
	timer := time.NewTicker(rateLimit)

	// Output channel to receive data back from the
	// routines.
	output := make(chan Output, 20)

	// For each device, spawn a goroutine
	//
	// In real-world, you would chunk this data up in a way
	// where you could spawn a predefined number of routines.
	for _, device := range devices {
		go func(d string) {
			Post(d, timer, output)
		}(device)
	}

	defer close(output)

	// Range through for the length of the device list.
	// Block until there is something to receive out of the channel,
	// then output the data received.
	for i := 1; i <= len(devices); i++ {
		p := <-output
		if p.Err != nil {

			logrus.Error(p.Err)
			continue
		}
		// logrus.Infof("%s:%d: received output from buffer", time.Now().UTC(), p.Data.Result[0].ID)
		fmt.Println(p.Data.Result)
	}
}

// Resp is a type that we expect to receive back.
// This type will be used to unmarshal the data received
// from the API
type Resp struct {
	Result []struct {
		Model string `json:"model"`
		ID    int    `json:"number"`
	} `json:"result"`
}

// Err is a type that we expect to receive back when we experience an error.
type Err struct {
	Status  int    `json:"status_code"`
	Error   string `json:"error_message"`
	Request string `json:"request"`
}

// Output type defines the data we expect to receive back from the
// routines. This type can be used as a data channel
type Output struct {
	Data Resp
	Err  error
}

// Post contains the actual http query we are making out
func Post(device string, timer *time.Ticker, output chan Output) {

	// Here, we wait in a blocking state until the timer func signals on its
	// channel with the current time.
	<-timer.C

	const query = `{"some_json_payload": "some_key"}`

	// Make a io.Reader type using a query payload
	payload := strings.NewReader(fmt.Sprintf(query, device))

	// NewRequest returns a new Request given a method, URL, and optional body.
	req, err := http.NewRequest(http.MethodPost, "https://test.service-now.com/", payload)
	if err != nil {
		output <- Output{Err: err}
		return
	}
	// Setting the request header for token auth
	req.Header.Set("API-KEY", os.Getenv("TOKEN"))

	res, err := http.DefaultClient.Do(req)
	if res != nil {
		// Note: We use a func to error check defer as opposed to using
		// defer res.Body.Close(), which will never return an error.
		defer func() {
			if defErr := res.Body.Close(); defErr != nil {
				err = errors.Wrap(defErr)
			}
		}()
	}
	if err != nil {
		output <- Output{Err: err}
		return
	}

	// Here, we are just checking to see if the status code is something larger than
	// we would expect to get back. In this example, I use 204, but there
	// are multiple ways to do this.
	if res.StatusCode > http.StatusNoContent {
		var body Err
		err = json.NewDecoder(res.Body).Decode(&body)
		output <- Output{Err: fmt.Errorf("%s:%d:%s:%s", device, res.StatusCode, res.Status, body.Error)}
		return
	}

	var response = new(Resp)
	if err = json.NewDecoder(res.Body).Decode(&[]interface{}{response}); err != nil {
		output <- Output{Err: err}
		return
	}
	// logrus.Infof("%s:%s: sending output to buffer", time.Now().UTC(), device)
	output <- Output{Data: *response}
}

// parseCSV function will take a CSV filename and attempt to parse the data into a list of
// strings.
func parseCSV() []string {
	header := true
	csvFile, err := os.Open("example.csv")
	if err != nil {
		// Don't panic! Handle the error properly.
		panic(err)
	}

	reader := csv.NewReader(bufio.NewReader(csvFile))
	var devices []string
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		// Trim out the header
		if header {
			header = false
			continue
		}
		devices = append(devices, strings.TrimSpace(line[0]))
	}
	return devices
}
