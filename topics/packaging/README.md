## Packaging


## Go Modules 
Go modules was introduced in Go v1.11 and is now the official dependency 
management system supported by the Go language team. 

Prior to releasing Go modules, dependencies were managed mostly manually, by pulling
down packages with the `go get` command syntax. The struggles associated with this
type of dependency management usually centered around inconsistent version types between
systems (developers having different versions of a package depending on when they last updated), 
and package maintainers deploying _sometimes_ breaking changes without any way to 
determine breaking change updates to minor updates.

Other dependency management systems were created prior to go modules (dep / glide for example)
However, most dependency management systems now are becoming depricated in favor of go modules.

##### Module Workflows
Below are some go-to commands when interacting with the module dependency system.
* `go mod init` Creates a new module, initializing the go.mod file that describes it.

* `go list -m all` Prints the current module’s dependencies.

* `go get` Fetch changes the required version of a dependency (or adds a new dependency).

* `go mod tidy` Removes unused dependencies.

* `go clean -modcache` Remove the entire module download cache, including 
unpacked source code of versioned dependencies.


#### References
+ [Using Go Modules](https://blog.golang.org/using-go-modules)
+ [Modules Wiki](https://github.com/golang/go/wiki/Modules)
+ [Modules Release Notes](https://golang.org/doc/go1.13#modules)
