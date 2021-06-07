package main

import (
	"fmt"
	"time"
)

func main() {
	// Create an unbuffered channel
	baton := make(chan int)

	// First runner to his mark
	go Runner(baton)

	// Start the race
	baton <- 1

	// Give the runners time to race
	time.Sleep(5 * time.Second)
}

func Runner(baton chan int) {
	var newRunner int

	// Wait to receive the baton
	runner := <-baton

	// Start running around the track
	fmt.Printf("%s - runner-%d running...\n", time.Now().String(), runner)

	// New runner to the line
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner %d To The Line\n", newRunner)
		go Runner(baton)
	}

	// Running around the track
	time.Sleep(1 * time.Second)

	// Is the race over
	if runner == 4 {
		fmt.Printf("Runner %d Finished, Race Over\n", runner)
		return
	}

	// Exchange the baton for the next runner
	fmt.Printf("Runner %d Exchange With Runner %d\n", runner, newRunner)
	baton <- newRunner
}
