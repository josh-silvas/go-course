// Pointers Example: In this example, we are going to initialize a blank structure, then
//   pass the structure to another function by address that will modify the data, then
//   print the struct out to terminal
//
// Before you begin:
//   1. Change to the practice directory
//       - cd practice/pointers
//   2. Complete the `TO DO` sections of the code
//   3. Check answers in answer/main.go

package main

import (
	"encoding/json"
	"log"
	"strings"
)

// DeviceData struct is a small example of data that we might want to know
// about a device.
type DeviceData struct {
	DeviceID  int    `json:"number"`
	Account   string `json:"account"`
	Hostname  string `json:"hostname"`
	IPAddress string `json:"ip_address"`
}

func main() {
	// TODO: 1. Initialize the DeviceData type with it's empty value.

	// TODO: 2. Call the modifyData function passing in the address (pointer) of the newly created DeviceData.

	// TODO: 4. Print out the structure you created. This should have the modified data from the modifyData func

}

func modifyData(d *DeviceData) {
	// TODO: 3. Use the res (response data) variable that contains device information to populate the *d (DeviceData)
	res := getData()

}

// getData is a small function that takes a json string and decodes it into a struct.
// This function is intended to mock the type of decoding and struct population that
// you might use when interfacing with a real API.
func getData() (d DeviceData) {
	apiResponse := `{
		"number": 879665,
		"account": "Golang.org Production",
		"hostname": "fw1.golang.org",
		"ip_address": "186.75.3.9"
	}`
	if err := json.NewDecoder(strings.NewReader(apiResponse)).Decode(&d); err != nil {
		log.Fatal(err)
	}
	return
}
