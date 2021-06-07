// Methods Answer: In this example, we will be defining an IP address structure and attaching a
//   CIDR method to the type, that will print the CIDR value of an IP/Netmask combination.
//
// Before you begin:
//   1. Change to the practice directory
//       - cd practice/methods
//   2. Complete the `TO DO` sections of the code
//   3. Check answers in answer/main.go

package main

import (
	"fmt"
	"net"
)

// Address struct defines the IP and Netmask types of an address
// TODO: 1. Define the struct with data types for 'IP' and 'Netmask' values.
type Address struct {
	IP      string
	Netmask string
}

func main() {
	// TODO: 2. Initialize an Address data type with a populated 'IP' and 'Netmask' values.
	var myStruct = Address{
		IP:      "50.56.20.4",
		Netmask: "255.255.255.240",
	}

	// TODO: 3. Print the IP address and Netmask values.
	fmt.Printf("Addr: %s\nMask: %s\n\n", myStruct.IP, myStruct.Netmask)

	// TODO: 5. Call the CIDR method using the Address receiver
	myStruct.CIDR()
}

// CIDR is a method attached to the Address data type that will convert the IP address and subnet mask into a
// CIDR value.
// TODO: 4. Convert this function into a method for the Address data type.
func (a Address) CIDR() {
	// Get the subnet mask bit size. Note, Size() returns the number of leading bits
	// in the netmask, along with the total bits of the address. For this address, the
	// return value will be 28 and 32, as 28 is the netmask in bits and 32 is the address
	// total size in bits. We will throw away the total bits in this case since we don't need them.
	leadingOnes, _ := net.IPMask(net.ParseIP(a.Netmask).To4()).Size()

	// Print the IP and netmask bit length. Ideally, we could do anything here on the CIDR value.
	// The idea is that the method is attached to the data structure of Address.
	fmt.Printf("CIDR: %s/%d\n", a.IP, leadingOnes)
}

// Results:
// ---------------------------------------------------------------------
// Addr: 50.56.20.4
// Mask: 255.255.255.240
//
// CIDR: 50.56.20.4/28
