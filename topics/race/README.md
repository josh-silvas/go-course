## Data Race Detector
Go has a built-in race detection tool that has been part of the language since Go1.1. Anytime 
you are using concurrent programming (goroutines), there is a chance for a data race to happen.

A data race can occur when two goroutines access the same memory location concurrently and at least one 
of the accesses is a write, so basically, reading/writing to the same memory location at the same time. 

The race detector will utilize the compiler to prepend all memory accesses with code that records when and how 
the memory was accessed. The runtime library will then watch for unsynchronized accesses to shared variables.

 - While using the race-detector, memory usage may increase by 5-10x and execution time by 2-20x.
 - Write load, integration, or unit tests that will accurately test a data-race. 

### How to Use
```aidl
go test -race pkg        // to test the package
go run -race example.go  // to run the source file
go build -race           // to build the command
go install -race pkg     // to install the package
```

#### Reference
  - [Introducing the Race Detector](https://blog.golang.org/race-detector)
  - [Detecting Race Conditions with Go](https://www.ardanlabs.com/blog/2013/09/detecting-race-conditions-with-go.html)
  - [How ThreadSanitizerAlgorithm](https://github.com/google/sanitizers/wiki/ThreadSanitizerAlgorithm)
