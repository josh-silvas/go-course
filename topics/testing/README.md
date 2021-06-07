## Testing in Go
The tests that you write for your Go program should be very close, if not identical, to the types of interactions that
a consumer of your program would expect. I typically do not try to do anything fancy in testing, just interface with the
program API and evaluate it's response to ensure that it's what a consumer would expect. 

The syntax deliberately avoids the use of assertions and leaves the responsibility up to the developer on how
to handle assertions and comparisons.
    
#### Writing Unit Tests
There isn't a lot about test that isn't covered in the standard language, so there isn't a large amount of new
information when writing tests. There are however, a couple of things outlined below.
 - When possible, use the `pkgname_test` package name for your test files.
 - Keep your tests close to the actual source code.

To write a test, create a function and import the `testing.T` type from the
testing package. For example `TestGetUser(t *testing.T)`:
>T is a type passed to Test functions to manage test state and support formatted test logs.
 Logs are accumulated during execution and dumped to standard output when done.
 A test ends when its Test function returns or calls any of the methods
 FailNow, Fatal, Fatalf, SkipNow, Skip, or Skipf. 

```go
func TestEqualFold(t *testing.T) {
	if !strings.EqualFold("test", "TeSt") {
		t.Fatal("unexpected response from EqualFold")
	}
	t.Logf("EqualFold assesed correctly")
}

=== RUN   TestEqualFold
--- PASS: TestEqualFold (0.00s)
main_test.go:71: EqualFold assesed correctly
PASS
```

#### Coverage
Go has some neat builtin tools to evaluate your test coverage for a given package. 
The syntax for generating a coverage profile is `go test -coverprofile <filename>`.
```go
go test -coverprofile cover.out
```

Once you have a coverage profile, you can optionally convert your coverage profile to
a more readable out put using the Go builtin tool `go tool cover`, for example:
```go
go tool cover -html=cover.out
```

#### Writing Example Tests
Examples within your `_test.go` file can serve a dual purpose of both testing a 
piece of code, as well as serving as example documentation in your API docs.

Writing an example test:
```go
func ExampleEqualFold() {
	fmt.Println(strings.EqualFold("test", "TeSt"))
	// Output:
	// true
}

=== RUN   ExampleEqualFold
--- PASS: ExampleEqualFold (0.00s)
PASS
```

#### Writing Benchmark Tests
Benchmarking is another form of testing that is built right into Go as a natively 
supported tool.

Not covered in this course is [Profiling Go Programs](https://blog.golang.org/profiling-go-programs).

>Benchmark functions are run several times by the testing package. 
The value of b.N will increase each time until the benchmark runner 
is satisfied with the stability of the benchmark.
>
>Any benchmark should be careful to avoid compiler optimisations eliminating 
the function under test and artificially lowering the run time of the benchmark.

```go
var result int

func BenchmarkFibComplete(b *testing.B) {
        var r int
        for n := 0; n < b.N; n++ {
                // always record the result of Fib to prevent
                // the compiler eliminating the function call.
                r = Fib(10)
        }
        // always store the result to a package level variable
        // so the compiler cannot eliminate the Benchmark itself.
        result = r
}
```
>[Reference](https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go)
>
#### References
+ [Writing Unit Test](https://blog.alexellis.io/golang-writing-unit-tests/)
+ [How to Write Benchmarks in Go](https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go)
