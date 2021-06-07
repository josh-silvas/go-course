# Go Training
This material is designed to be used as a 2 to 3 day class to go over various 
mechanisms within the language as well as 3rd party packages that are useful 
in some of our practical application labs.

## Prior to Class
In the interest of time, there are some basic topics that are not covered in this course and 
are assumed knowledge prior to taking this class. A good primer would be to review these resources:
 - [Golang Tour#Basics](https://tour.golang.org/basics/1)
 - [Go By Example](https://gobyexample.com/)
 
You can also find a large list of books available and the list is maintained by the Golang team. 
This site should have the most up-to-date information on learning resources:
 - [Golang Books](https://github.com/golang/go/wiki/Books)

## Setting Up Your Environment
This course info will assume that you have already gone through the tasks to 
get a Go environment setup locally on your computer. 

Follow the setup instructions [HERE](https://github.com/josh5276/go-course/tree/master/SETUP.md)
to setup your environment.

## Some Go Key Points
Some key definitions about the language that may be different from your current
language of choice

 + Golang uses a static [type system](https://en.wikipedia.org/wiki/Type_system)
   + Static typing allows more opportunities for compiler optimizations
   + Finding type-mismatching at compile time causes less code errors
   + Other languages do this already  (.NET, C, Java), however:
     + Go's implementation of polymorphism allows for static typing to feel dynamic by matching interface signatures
     
 + Golang focuses on concurrency
   + One of the challenges that Go set out to solve is scalability of large applications. Golang can scale at 
   large using its built-in concurrency design.
   + Often in other languages, an effective concurrency model usually comes at the cost of ease-of-use. Go
   is able to achieve efficient concurrency and it's actually pretty easy to implement. 
 
 + Ease of Use
   + Believe it or not, most of the fundamentals in Go can be taught over a few hours
   in a classroom setting. There is continued development around best practices and 
   code organization/optimization that will need to come with this as well, but the core
   language is fairly simple.
   
 + Compiling and Distribution
   + Go compiles to a single native binary (as opposed to an interpreter/VM), which makes deploying 
   an application written in Go as easy as copying the application file to 
   the destination.

## The Topics
This section contains all of the topics covered in this course. The topics are arranged in an order that
I think is the most comprehensive from a learning perspective in class, however, everyone learns differently,
and you are more than welcome to step through these topics in any order that works for you.

##### Fundamentals
 - [Variables/Structs](https://github.com/josh5276/go-course/tree/master/topics/variables_structs)
   ([Practice](https://github.com/josh5276/go-course/blob/master/practice/variables_structs/main.go) /
    [Answer](https://github.com/josh5276/go-course/blob/master/practice/variables_structs/answer/main.go))
 - [Slices](https://github.com/josh5276/go-course/tree/master/topics/slices)
   ([Practice](https://github.com/josh5276/go-course/blob/master/practice/slices/main.go) /
    [Answer](https://github.com/josh5276/go-course/blob/master/practice/slices/answer/main.go))
 - [Maps](https://github.com/josh5276/go-course/tree/master/topics/maps)
   ([Practice](https://github.com/josh5276/go-course/blob/master/practice/maps/main.go) /
    [Answer](https://github.com/josh5276/go-course/blob/master/practice/maps/answer/main.go))
 - [Functions](https://github.com/josh5276/go-course/tree/master/topics/functions)
   ([Practice](https://github.com/josh5276/go-course/blob/master/practice/functions/main.go) /
    [Answer](https://github.com/josh5276/go-course/blob/master/practice/functions/answer/main.go))
 - [Pointers](https://github.com/josh5276/go-course/tree/master/topics/pointers)
   ([Practice](https://github.com/josh5276/go-course/blob/master/practice/pointers/main.go) /
    [Answer](https://github.com/josh5276/go-course/blob/master/practice/pointers/answer/main.go))
 - [Methods](https://github.com/josh5276/go-course/tree/master/topics/methods)
   ([Practice](https://github.com/josh5276/go-course/blob/master/practice/methods/main.go) /
    [Answer](https://github.com/josh5276/go-course/blob/master/practice/methods/answer/main.go))
 - [Packaging](https://github.com/josh5276/go-course/tree/master/topics/packaging)
    ([Pkg Example](https://github.com/josh5276/go-course/blob/master/practice/packaging/01-example-pkg/main.go) /
     [Mod Practice](https://github.com/josh5276/go-course/blob/master/practice/packaging/02-practice-modules/main.go))

##### Interfaces
 - [Interfaces](https://github.com/josh5276/go-course/tree/master/topics/interfaces)
   ([Practice](https://github.com/josh5276/go-course/blob/master/practice/interfaces/main.go) /
    [Answer](https://github.com/josh5276/go-course/blob/master/practice/interfaces/answer/main.go))
   - [Example](https://github.com/josh5276/go-course/blob/master/practice/interfaces/example/main.go)
 
##### Error Handling
 - [Standard Errors](https://github.com/josh5276/go-course/tree/master/topics/error)
   ([Practice](https://github.com/josh5276/go-course/blob/master/practice/errors/main.go) /
    [Answer](https://github.com/josh5276/go-course/blob/master/practice/errors/answer/main.go))
 - [Custom Errors](https://github.com/josh5276/go-course/tree/master/topics/error_custom)

##### Concurrent Programming
 - [Routines](https://github.com/josh5276/go-course/tree/master/topics/routines)
   ([Practice](https://github.com/josh5276/go-course/blob/master/practice/concurrency/main.go) /
    [Answer](https://github.com/josh5276/go-course/blob/master/practice/concurrency/answer/main.go))
 - [Channels](https://github.com/josh5276/go-course/tree/master/topics/channels)
   - Examples: ([UnBuffered](https://github.com/josh5276/go-course/blob/master/practice/channels/01-example-unbuffered/main.go) / 
    [Buffered](https://github.com/josh5276/go-course/blob/master/practice/channels/01-example-buffered/main.go) / 
    [Real-World](https://github.com/josh5276/go-course/blob/master/practice/channels/01-example-realworld/main.go))

###### Hands On 
 * [Working with APIs](https://github.com/josh5276/go-course/blob/master/practice/api/main.go)
 
##### Go Tools
 - [Testing](https://github.com/josh5276/go-course/tree/master/topics/testing)
   - [UnitTest Practice](https://github.com/josh5276/go-course/blob/master/practice/testing_unit/main_test.go)
   - [HTTPTest Example](https://github.com/josh5276/go-course/blob/master/practice/testing_http/example/main_test.go)
 - [Race Detector](https://github.com/josh5276/go-course/tree/master/topics/race)
   - [Example](https://github.com/josh5276/go-course/blob/master/practice/race/)

## Extras
 - [GolangCI-Lint](https://github.com/golangci/golangci-lint#golangci-lint)
 - [Goreleaser](https://goreleaser.com/)

## Join the Gopher Community!
##### Gopher Slack Instance
Join the Slack community of Gophers by registering here:
1. Fill out your name and email address [HERE](https://gophersinvite.herokuapp.com/)
2. Check your email, and follow the link to the slack application.

 