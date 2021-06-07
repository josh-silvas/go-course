package main

import (
	"fmt"
	"math/rand"
	"time"
)

var bufferSize = 20

func main() {
	// Example of signaling with data
	ch := make(chan string, bufferSize)

	// Range through the bufferSize and spawn a routine for each that will
	// sleep at a random interval, then send a string into the channel.
	for i := 1; i <= bufferSize; i++ {
		go func(i int) {
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			ch <- fmt.Sprintf("send %d at: %s", i, time.Now().UTC())
		}(i)
	}

	for i := 1; i <= bufferSize; i++ {
		// Receive data from the channel when a signal is sent from
		// one of the goroutines
		fmt.Println(<-ch)
	}
}
