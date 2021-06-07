## Concurrency
Concurrency is essentially taking a set of computational tasks and breaking them up to run in any particular
order to run across multiple execution paths, as opposed to serially on a processor.

This is one of Go's strong suits as a result of ease of use and behind-the-curtain scheduling. It's 
important to note that there are a lot of cases where concurrency does not actually solve a problem.
For concurrency to make sense, a problem where concurrency is the answer should be completely obvious and 
add value to the application, i.e., enough of a performance gain for the complexity cost.

Some pointers on concurrency:
1. Every one of your coroutines should terminate before the return of main().

2. Goroutines are very cheap to run, and you can launch hundreds of thousands of them.
Your system however, does have a limit. Understand the load that you could introduce onto your
system or your network. Rate limiting might be a consideration.

3. Channels and mutexes can create backpressure when goroutines are required to wait.
Watch the backpressure in your application. Backpressure happens when your application
receives too much work, and starts affecting the performance. Timeouts and closes for channels are 
good examples of ways to handle backpressure.
 

### Scheduling
Before we get into the concurrency, lets take a look at the way scheduling tasks against a set of
machine processors and cores work with Go's runtime.

##### The OS Scheduler
Accessing data from main memory has such a high latency cost 
(~100 to ~300 clock cycles) that processors and cores have local caches to
keep data close to the hardware threads that need it. Accessing data from 
caches have a much lower cost (~3 to ~40 clock cycles) depending on the 
cache being accessed.

![Example memory cache](https://github.com/josh5276/go-course/blob/master/common/img/example_cache.png)

Your OS schedulers job is to take threads that are able to run and execute on those threads
in a manner that allows the system to be utilized to it's fullest potential, while
prioritizing threads that need to be run.
 - The OS scheduler can be susceptible to false sharing
  - > If one Thread on a given core makes a change to its copy of the cache line,
  then through the magic of hardware, all other copies of the same cache 
  line have to be marked dirty. When a Thread attempts read or write access 
  to a dirty cache line, main memory access (~100 to ~300 clock cycles) is 
  required to get a new copy of the cache line.
  -- [Ref: Scheduling in Go](https://www.ardanlabs.com/blog/2018/08/scheduling-in-go-part1.html)

##### The Go Runtime Scheduler
Sits in the user-space on the OS, not the kernel space. So 

The Go scheduler takes the following items into account for system and code scheduling
 - Number of virtual cores 
 - One OS thread is assigned for each virtual core
 - Each Goroutine will get it's own assignment
 - Global Run Queue
   - Where Goroutines sit and wait for a virtual core assignment
 - Local Run Queue
    - One per-virtual core. Queues Goroutines scheduled to be executed within 
    the cores context.

 Any given Goroutine can be in one of three states:
  - Waiting:   Stopped and not running (mutex/system-level/networkIO calls)
  - Runnable:  Able to run, but waiting for an assignment to an OS thread
  - Executing: Currently running

When does the scheduler make scheduling decisions?
 - The use of the keyword go
   - How to spawn a goroutine in code
 - Garbage collection
   - The garbage collector spawns it's own set of goroutines when it's running
 - System calls
   - System calls can be blocking calls. 
 - Synchronization and Orchestration
   - Mutexes and channel operations
    
Network pollers give the OS the opportunity to place a routine waiting for a 
network operation in it's own queue so that other routines can be processed 
without blocking a thread. 

Once the network operation completes, the queued routine is placed back into 
the local run queue in the back of the line.

When a call cannot be switched to the network poller, the go scheduler will 
take the virtual representation of the OS thread, and move it off of the 
virtual core, allowing another thread to process on that core while the other
 routine is in a blocking state.

The scheduler also has the ability to take routines from a local run queue 
of another thread, or from the global run queue, to process them if it does 
not have any work to do.

The GO schedulers entire intent is to massage routines that need processing 
and place them on the right queue at the right time, so that from the OS 
level scheduling, it appears as though there is never a time when we are 
blocking a process from running. 

![Schedulers together](https://github.com/josh5276/go-course/blob/master/common/img/OS_Go_SchedulerMap.png)

1. M's will execute a routine that is ready to be run
2. Asynchronous network calls can be placed on the network poller to free up the processor
    2. When the async network call returns, it will be placed on the back of the local run queue
3. Synchronous blocking calls will detach the M0 from the processor, so that a new Mn can continue to process.
    3. Once the block call is complete, the routine goes back into a local run queue, and M0 is saved 
    for future use.
4. If M1 does not have work, it can steal routines from another local run queue
5. If M1 does not have work, and there are no local queued routines in another queue, M1 will pull a
 routine from the global run queue.
 
 ### Example
  - [Source Example](https://github.com/josh5276/go-course/blob/master/topics/routines/example.go)
  
 #### Reference
  - [Concurrency is Not Parallelism - Rob Pike](https://www.youtube.com/watch?v=cN_DpYBzKso)