// Maps Answer: In this example, we are providing a problem that could be solved with a
// map implementation.
// Before you begin:
//   1. Change to the practice directory
//       - cd practice/maps
//   2. Complete the `TO DO` sections of the code
//   3. Check answers in answer/main.go

package main

import (
	"fmt"
	"sort"
)

// myMap is a collection of key:value pairs that contains a list of the different
// platform.model's returned, with a value of a standard platform-type
// classification.
var myMap = map[string]string{
	"Load-Balancer":                             "Load-Balancer",
	"Cisco Pix 501 (10 IP)":                     "Firewall - Cisco PIX",
	"BigIP- F5 1600":                            "Load-Balancer",
	"Cisco ASA":                                 "Firewall - Cisco ASA",
	"Cisco ASA 5505 Sec+":                       "Firewall - Cisco ASA",
	"Firewall - Cisco PIX":                      "Firewall - Cisco PIX",
	"Firewall - Juniper SRX":                    "Firewall - Juniper SRX",
	"Cisco ASA 5505 Unlimited (Orange Sticker)": "Firewall - Cisco ASA",
	"Cisco ASA 5508 X":                          "Firewall - Cisco ASA",
	"Cisco ASA 5508-X with SSD":                 "Firewall - Cisco ASA",
	"Cisco ASA 5510 Sec+":                       "Firewall - Cisco ASA",
	"Cisco ASA 5515":                            "Firewall - Cisco ASA",
	"Cisco Pix 501 (50 IP)":                     "Firewall - Cisco PIX",
	"Cisco Pix 515URLL":                         "Firewall - Cisco PIX",
	"Firewall - Cisco":                          "Firewall - Cisco",
	"Firewall - Cisco ASA":                      "Firewall - Cisco ASA",
	"Firewall - Palo Alto":                      "Firewall - Palo Alto",
	"Brocade ADX 1000 Series":                   "Load-Balancer",
	"Load Balancer (Virtual) - VMWare":          "Load-Balancer",
	"Cisco Pix 506":                             "Firewall - Cisco PIX",
	"Cisco Pix 515R":                            "Firewall - Cisco PIX",
}

func main() {
	// TODO: 1. Add a key:value of "Cisco ASA 5525": "Firewall - Cisco ASA"
	myMap["Cisco ASA 5525"] = "Firewall - Cisco ASA"

	// TODO: 2. PIX's are being removed. Delete any key with a value of "Firewall - Cisco PIX"
	for key, value := range myMap {
		if value == "Firewall - Cisco PIX" {
			delete(myMap, key)
		}
	}

	// TODO: 3. Sort the keys alphabetically
	var keys []string
	for k := range myMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// TODO: 4. Print out the new, modified and sorted map.
	for _, key := range keys {
		fmt.Printf("%-45s%s\n", key, myMap[key])
	}
}

// Results:
// ---------------------------------------------------------------------
// BigIP- F5 1600                               Load-Balancer
// Brocade ADX 1000 Series                      Load-Balancer
// Cisco ASA                                    Firewall - Cisco ASA
// Cisco ASA 5505 Sec+                          Firewall - Cisco ASA
// Cisco ASA 5505 Unlimited (Orange Sticker)    Firewall - Cisco ASA
// Cisco ASA 5508 X                             Firewall - Cisco ASA
// Cisco ASA 5508-X with SSD                    Firewall - Cisco ASA
// Cisco ASA 5510 Sec+                          Firewall - Cisco ASA
// Cisco ASA 5515                               Firewall - Cisco ASA
// Cisco ASA 5525                               Firewall - Cisco ASA
// Firewall - Cisco                             Firewall - Cisco
// Firewall - Cisco ASA                         Firewall - Cisco ASA
// Firewall - Juniper SRX                       Firewall - Juniper SRX
// Firewall - Palo Alto                         Firewall - Palo Alto
// Load Balancer (Virtual) - VMWare             Load-Balancer
// Load-Balancer                                Load-Balancer
