package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Setting a maxProcs value of 1 will only allow the scheduler to
// use 1 processor at a time. This means that even though we are creating
// multiple routines, they will still get processed serially.
//
// Setting to 2 will allow the scheduler to process these routines
// in parallel.
const maxProcs = 2

func init() {
	runtime.GOMAXPROCS(maxProcs)
}

func main() {
	fmt.Printf("MaxProcs Allowed: %d\n", maxProcs)
	fmt.Printf("NumCPUs: %d\n", runtime.NumCPU())

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Kicking off routines")
	go routineOne(&wg)
	go routineTwo(&wg)

	wg.Wait()

	fmt.Println("We done!")
}

func routineOne(wg *sync.WaitGroup) {
	for r := 0; r <= 30; r++ {
		fmt.Printf("\033[31mRoutineONE\033[39m\n")
	}
	wg.Done()
}

func routineTwo(wg *sync.WaitGroup) {
	for r := 0; r <= 30; r++ {
		fmt.Printf("\033[34mRoutineTWO\033[39m\n")
	}
	wg.Done()
}
