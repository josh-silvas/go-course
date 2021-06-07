// Slices Answer: In this example, we are interacting with a pre-created slice of strings
//
// Before you begin:
//   1. Change to the practice directory
//       - cd practice/slices
//   2. Complete the `TO DO` sections of the code
//   3. Check answers in answer/main.go

package main

import (
	"fmt"
	"strings"
)

var mySlice = []string{
	"Load-Balancer",
	"Cisco Pix 501 (10 IP)",
	"BigIP- F5 1600",
	"Cisco ASA",
	"Cisco ASA 5505 Sec+",
	"Firewall - Cisco PIX",
	"Firewall - Juniper SRX",
	"Cisco ASA 5505 Unlimited (Orange Sticker)",
	"Cisco ASA 5508 X",
	"Cisco ASA 5508-X with SSD",
	"Cisco ASA 5510 Sec+",
	"Cisco ASA 5515",
	"Cisco Pix 501 (50 IP)",
	"Cisco Pix 515URLL",
	"Firewall - Cisco",
	"Firewall - Cisco ASA",
	"Firewall - Palo Alto",
	"Brocade ADX 1000 Series",
	"Load Balancer (Virtual) - VMWare",
	"Cisco Pix 506",
	"Cisco Pix 515R",
}

func main() {
	// TODO: 1. Print out the length (len) and capacity (cap) of the current slice mySlice
	fmt.Printf("Length Before: %d\n", len(mySlice))
	fmt.Printf("Capaci Before: %d\n\n", cap(mySlice))

	// TODO: 2. Add a new NetDevice type "Palo Alto" to the existing slice
	mySlice = append(mySlice, "Palo Alto")

	// TODO: 3. Print the last 5 items in the slice
	fmt.Print("Last 5 items in slice:\n")
	fmt.Println(strings.Repeat("-", 25))
	fmt.Printf("%s\n\n", strings.Join(mySlice[len(mySlice)-5:], "\n"))

	// TODO: 4. Print out the new length and capacity after adding a new value
	fmt.Printf("Length After: %d\n", len(mySlice))
	fmt.Printf("Capaci After: %d\n", cap(mySlice))

}

// Results:
// ---------------------------------------------------------------------
// Length Before: 21
// Capaci Before: 21
//
// Last 5 items in slice:
// -------------------------
// Brocade ADX 1000 Series
// Load Balancer (Virtual) - VMWare
// Cisco Pix 506
// Cisco Pix 515R
// Palo Alto
//
// Length After: 22
// Capaci After: 44
