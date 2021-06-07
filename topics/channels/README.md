## Channels
A channel provide a way to share memory resource across multiple concurrent
routines running while maintaining the integrity of that data stored in 
memory.

Channels serve as a way to prevent data races in concurrent programming. Data
races can happen when:
 - A routine is writing data to a location in memory, while another 
 routine is also reading data from the same location in memory. 
 - Or, when a routine is writing data to a location in memory at the same
  time that another routine is also writing data to the same location in memory.
  
##### Send Operator
`chan <- "expression"` 
A send statement sends a value on a channel. 
 - The channel expression must be of channel type
 - The channel direction must permit send operations
 - The type of the value to be sent must be assignable to the channel's element type
 
##### Receive Operator
`"value" <- chan` The value received from the channel ch. 
 - The channel direction must permit receive operations
 - The type of the receive operation is the element type of the channel.
 - This will block until a value is available. 
 - If you are receiving from a nil channel, it blocks forever. 

A receive operation on a closed channel can always proceed immediately, 
yielding the element type's zero value after any previously sent values 
have been received. 

### Considerations
When creating a channel for your application, consider these questions. These control knobs of a channel
will define its behavior and give you mechanisms to alter the way your code behaves between programs.
1. Buffered or Unbuffered
2. State of the channel
3. With or without data
--------
##### Unbuffered Channels
Guaranteed Delivery.
Unbuffered channels are ones that are created with no capacity to 
send/receive data. This channel type will allow for complete synchronization between
routines. That means when there is a value to be passed to a channel,
the sending function will block and wait until there is something to 
receive the data from the channel.  

Likewise, if a routine is not readily available to send data on an unbuffered
channel, the receiving routine will wait until there is data to be received.

##### Buffered Channels
No Guarantee of Delivery.
Buffered channels are created with a `make` statement, containing a value equal to the size of the 
buffered channel. 
 * `ch := make(chan int, 50)`

Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty. 

--------
##### State
 * **Nil State (x -- x)**: When a channel is in a nil state, any send or receive attempted on the channel will block
    * `var ch chan int`
 * **Open State (√ -- √)**: When a channel is in an open state, signals can be sent and received.
    * `ch := make(chan int)`
 * **Closed State (! -- √)**: When a channel is placed into a closed state, signals can no longer be sent but it’s still 
 possible to receive signals.
    * `close(ch)`
 
--------
##### With Data
An example of signaling with data would be: `chan <- "expression"`
   * A goroutine is being asked to start a new task.
   * A goroutine reports back a result.
   
##### Without Data
Signaling without data can be in the form of: `close(ch)` _Mostly used for cancellation._
   * A goroutine is being told to stop what they are doing.
   * A goroutine reports back they are done with no result.
   * A goroutine reports that it has completed processing and shut down.
---------

#### Reference

 - [Go Spec (Channel Types)](https://golang.org/ref/spec#Channel_types)
 - [Effective Go (Channels)](https://golang.org/doc/effective_go.html#channels)
 - [Go101 Channels](https://go101.org/article/channel.html)