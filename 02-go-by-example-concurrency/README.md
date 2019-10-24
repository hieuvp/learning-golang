# Go by Example - Concurrency


## Table of Contents
<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->


- [Goroutines](#goroutines)
- [Channels](#channels)
- [Channel Buffering](#channel-buffering)
- [Channel Synchronization](#channel-synchronization)
- [Channel Directions](#channel-directions)
- [Select](#select)
- [Timeouts](#timeouts)
- [Non-Blocking Channel Operations](#non-blocking-channel-operations)
- [Closing Channels](#closing-channels)
- [Range over Channels](#range-over-channels)
- [Timers](#timers)
- [Tickers](#tickers)
- [Worker Pools](#worker-pools)
- [WaitGroups](#waitgroups)
- [Rate Limiting](#rate-limiting)
- [Atomic Counters](#atomic-counters)
- [Mutexes](#mutexes)
- [Stateful Goroutines](#stateful-goroutines)
- [Sorting](#sorting)
- [Sorting by Functions](#sorting-by-functions)
- [Panic](#panic)
- [Defer](#defer)
- [References](#references)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->


## Goroutines

> A **goroutine** is a lightweight thread of execution.

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=goroutines.go) -->
<!-- The below code snippet is automatically added from goroutines.go -->
```go
package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 4; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// Run "synchronously"
	f("Direct")

	// Use "go f(string)" to invoke this function in a "goroutine"
	// This new "goroutine" will execute concurrently with the calling one
	go f("Goroutine")

	// Start a "goroutine" for an "anonymous function" call
	go func(message string) {
		fmt.Println(message)
	}("Going")

	// Two function calls above are running "asynchronously" in separate "goroutines" now

	// Pause for at least the "duration"
	duration := time.Second
	fmt.Println("Sleep with Duration :", duration)
	time.Sleep(duration)
	fmt.Println("Done")
}
```
<!-- AUTO-GENERATED-CONTENT:END -->

```bash
$ go run goroutines.go

# Direct : 0
# Direct : 1
# Direct : 2
# Direct : 3

# Goroutine : 0
# Sleep with Duration : 1s
# Goroutine : 1
# Goroutine : 2
# Goroutine : 3
# Going

# Done
```


## Channels

> **Channels** are the pipes that connect concurrent goroutines.
> You can send values into **channels** from one goroutine and receive those values into another goroutine.

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=channels.go) -->
<!-- The below code snippet is automatically added from channels.go -->
```go
package main

import (
	"fmt"
	"reflect"
)

func main() {

	// Create a new channel with "make(chan value-type)"
	// Channels are typed by the values they convey
	messages := make(chan string)
	fmt.Println(reflect.TypeOf(messages))

	// "channel <-": send a value into a channel
	go func() { messages <- "ping" }()

	// "<-channel": receive a value from the channel
	msg := <-messages
	fmt.Println(msg)
}
```
<!-- AUTO-GENERATED-CONTENT:END -->

```bash
$ go run channels.go

# chan string
# ping
```


## Channel Buffering

- By default, **channels** are **unbuffered**, meaning that they will only accept sends (`chan <-`) if there is a corresponding receive (`<-chan`) ready to receive the sent value.
- **Buffered channels** accept a limited number of values without a corresponding receiver for those values.

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=channel-buffering.go) -->
<!-- The below code snippet is automatically added from channel-buffering.go -->
```go
package main

import (
	"fmt"
	"reflect"
)

func main() {

	// Here we make a channel of strings buffering up to "2" values
	messages := make(chan string, 2)
	fmt.Println(reflect.TypeOf(messages))

	// Because this channel is buffered,
	// we can send these values into the channel without a corresponding concurrent receive
	messages <- "buffered"
	messages <- "channel"

	// Later we can receive these "2" values as usual
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
```
<!-- AUTO-GENERATED-CONTENT:END -->

```bash
$ go run channel-buffering.go

# chan string

# buffered
# channel
```


## Channel Synchronization

> We can use **channels** to synchronize execution across **goroutines**.

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=channel-synchronization.go) -->
<!-- The below code snippet is automatically added from channel-synchronization.go -->
```go
package main

import (
	"fmt"
	"time"
)

// This is the function we will run in a goroutine
// The "done" channel will be used to notify another goroutine
func worker(done chan bool) {
	fmt.Print("Working... ")
	time.Sleep(time.Second)
	fmt.Println("Done")

	// Send a value to notify that we are done
	done <- true
}

func main() {

	// Start a "worker" goroutine
	done := make(chan bool, 1)
	go worker(done)

	// Block until we receive a notification from the "worker" on the channel
	<-done
}
```
<!-- AUTO-GENERATED-CONTENT:END -->

```bash
$ go run channel-synchronization.go

# Working... Done
```

- If you removed the `<-done` line from this program, the program would exit before the worker even started.


## Channel Directions

> When using **channels** as function parameters, you can specify if a **channel** is meant to only send or receive values.

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=channel-directions.go) -->
<!-- The below code snippet is automatically added from channel-directions.go -->
```go
package main

import "fmt"

// "pings chan<- string": a channel for sending values,
// would be a compile-time error to try to receive on this channel
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// "pings <-chan string": a channel for receives
// "pongs chan<- string": a channel for sends
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
```
<!-- AUTO-GENERATED-CONTENT:END -->

```bash
$ go run channel-directions.go

# passed message
```


## Select

> `select` lets you wait on multiple channel operations.

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=select.go) -->
<!-- The below code snippet is automatically added from select.go -->
```go
package main

import (
	"fmt"
	"time"
)

func main() {

	// We will "select" across two channels
	c1 := make(chan string)
	c2 := make(chan string)

	// Each channel will receive a value after some amount of time
	go func() {
		time.Sleep(4 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// Use "select" to await both of these values simultaneously
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Received :", msg1)
		case msg2 := <-c2:
			fmt.Println("Received :", msg2)
		}
	}
}
```
<!-- AUTO-GENERATED-CONTENT:END -->

```bash
$ time go run select.go

# Received : two
# Received : one

# go run select.go  0.29s user 0.29s system 13% cpu 4.474 total
```

- The total execution time is only `~4s` since both `4s Sleep` and `2s Sleep` execute concurrently.


## Timeouts

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=timeouts.go) -->
<!-- The below code snippet is automatically added from timeouts.go -->
```go
package main

import (
	"fmt"
	"time"
)

func main() {

	// The channel is "buffered", so the "send" in the goroutine is "non-blocking"
	// A common pattern to prevent goroutine leaks in case the channel is never "read"
	c1 := make(chan string, 1)

	// In contrast, "make(chan string)" is an "unbuffered channel"
	// The sender will block on the channel until the receiver receives the data from the channel

	// In both channels,
	// the receiver will always block on the channel until sender sends data into the channel

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result from c1"
	}()

	// Since "select" proceeds with the first receive that is ready,
	// we will take the timeout case if the operation takes more than the allowed "1s"
	select {
	case response := <-c1:
		fmt.Println(response)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout c1 after 1s")
	}

	// If we allow a longer timeout of "3s",
	// then the receive from "c2" will succeed
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result from c2"
	}()

	select {
	case response := <-c2:
		fmt.Println(response)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout c2 after 3s")
	}
}
```
<!-- AUTO-GENERATED-CONTENT:END -->

```bash
$ go run timeouts.go

# timeout c1 after 1s
# result from c2
```


## Non-Blocking Channel Operations

> Use `select` with a `default` clause to implement non-blocking sends, receives, and non-blocking multi-way selects.

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=non-blocking-channel-operations.go) -->
<!-- The below code snippet is automatically added from non-blocking-channel-operations.go -->
```go
package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// Here is a "non-blocking receive"
	// If a value is available on "messages" then "select" will take the "<-messages" case
	// If not it will "immediately" take the "default" case
	select {
	case message := <-messages:
		fmt.Println("received message :", message)
	default:
		fmt.Println("no message received")
	}

	// A "non-blocking send" works similarly
	// "message" cannot be sent to the "messages" channel,
	// because the channel has no buffer and there is no receiver
	// Therefore the "default" case is selected
	message := "hi"
	select {
	case messages <- message:
		fmt.Println("sent message :", message)
	default:
		fmt.Println("no message sent")
	}

	// Here we attempt "non-blocking receives" on both "messages" and "signals"
	select {
	case message := <-messages:
		fmt.Println("received message :", message)
	case signal := <-signals:
		fmt.Println("received signal :", signal)
	default:
		fmt.Println("no activity")
	}
}
```
<!-- AUTO-GENERATED-CONTENT:END -->

```bash
$ go run non-blocking-channel-operations.go

# no message received
# no message sent
# no activity
```


## Closing Channels

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=closing-channels.go) -->
<!-- The below code snippet is automatically added from closing-channels.go -->
```go
package main

import "fmt"

func main() {

	// Main goroutine
	jobs := make(chan int, 5)
	done := make(chan bool)

	// Worker goroutine
	go func() {

		// Repeatedly receives from "jobs"
		for {
			job, more := <-jobs

			// The "more" value will be "false",
			// if "jobs" has been closed and
			// all values in the channel have already been received
			if more {
				fmt.Println("received job :", job)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// Send "3" jobs to the "worker" over the "jobs" channel
	for job := 1; job <= 3; job++ {
		jobs <- job
		fmt.Println("sent job :", job)
	}

	// If a channel is closed, no values can be sent on it
	// "close" a channel to indicate "completion"
	close(jobs)
	fmt.Println("sent all jobs")

	// We await the "worker" using the synchronization approach
	<-done
}
```
<!-- AUTO-GENERATED-CONTENT:END -->

```bash
$ go run closing-channels.go

# sent job : 1
# sent job : 2
# sent job : 3

# sent all jobs

# received job : 1
# received job : 2
# received job : 3

# received all jobs
```


## Range over Channels

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=range-over-channels.go) -->
<!-- The below code snippet is automatically added from range-over-channels.go -->
```go
package main

import "fmt"

func main() {

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"

	// Here we "close" the channel,
	// the iteration will terminate after receiving "2" elements
	close(queue)

	// It is possible to "close" a non-empty channel but still have the remaining values be received
	// This "range" iterates over each element as it is received from "queue"
	for element := range queue {
		fmt.Println(element)
	}
}
```
<!-- AUTO-GENERATED-CONTENT:END -->

```bash
$ go run range-over-channels.go

# one
# two
```


## Timers


## Tickers


## Worker Pools


## WaitGroups


## Rate Limiting


## Atomic Counters


## Mutexes

> In the previous example we saw how to manage simple counter state using atomic operations.
> For more complex state we can use a mutex to safely access data across multiple goroutines.

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=mutexes.go) -->
<!-- The below code snippet is automatically added from mutexes.go -->
```go
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	// For our example the state will be a map
	var state = make(map[int]int)

	// This mutex will synchronize access to state
	var mutex = &sync.Mutex{}

	// We'll keep track of how many read and write operations we do
	var readOps uint64
	var writeOps uint64

	// Here we start 100 goroutines to execute repeated reads against the state,
	// once per millisecond in each goroutine
	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {

				// For each read we pick a key to access,
				// Lock() the mutex to ensure exclusive access to the state,
				// read the value at the chosen key,
				// Unlock() the mutex, and increment the readOps count
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&readOps, 1)

				// Wait a bit between reads
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// We'll also start 10 goroutines to simulate writes, using the same pattern we did for reads
	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Let the 10 goroutines work on the state and mutex for a second
	time.Sleep(time.Second)

	// Take and report final operation counts
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)

	// With a final lock of state, show how it ended up
	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}
```
<!-- AUTO-GENERATED-CONTENT:END -->

- Running the program shows that we executed about 90,000 total operations against our mutex-synchronized state.

```bash
$ go run mutexes.go
```

- Next we'll look at implementing this same state management task using only goroutines and channels.


## Stateful Goroutines

> In the previous example we used explicit locking with mutexes to synchronize access to shared state across multiple goroutines.
> Another option is to use the built-in synchronization features of goroutines and channels to achieve the same result.
> This channel-based approach aligns with Go’s ideas of sharing memory by communicating and having each piece of data owned by exactly 1 goroutine.

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=stateful-goroutines.go) -->
<!-- The below code snippet is automatically added from stateful-goroutines.go -->
```go
package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// In this example our state will be owned by a single goroutine
// This will guarantee that the data is never corrupted with concurrent access
// In order to read or write that state, other goroutines will send messages to the owning goroutine and receive corresponding replies
// These readOp and writeOp structs encapsulate those requests and a way for the owning goroutine to respond
type readOp struct {
	key  int
	resp chan int
}
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {

	// As before we'll count how many operations we perform
	var readOps uint64
	var writeOps uint64

	// The reads and writes channels will be used by other goroutines to issue read and write requests, respectively
	reads := make(chan readOp)
	writes := make(chan writeOp)

	// Here is the goroutine that owns the state,
	// which is a map as in the previous example but now private to the stateful goroutine
	// This goroutine repeatedly selects on the reads and writes channels,
	// responding to requests as they arrive
	// A response is executed by first performing the requested operation
	// and then sending a value on the response channel resp to indicate success (and the desired value in the case of reads)
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// This starts 100 goroutines to issue reads to the state-owning goroutine via the reads channel
	// Each read requires constructing a readOp, sending it over the reads channel,
	// and the receiving the result over the provided resp channel
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// We start 10 writes as well, using a similar approach
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Let the goroutines work for a second
	time.Sleep(time.Second)

	// Finally, capture and report the op counts
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}
```
<!-- AUTO-GENERATED-CONTENT:END -->

- Running our program shows that the goroutine-based state management example completes about 80,000 total operations.

```bash
$ go run stateful-goroutines.go
```

- For this particular case the goroutine-based approach was a bit more involved than the mutex-based one. It might be useful in certain cases though, for example where you have other channels involved or when managing multiple such mutexes would be error-prone. You should use whichever approach feels most natural, especially with respect to understanding the correctness of your program.


## Sorting

> Go's sort package implements sorting for builtins and user-defined types.
> We'll look at sorting for builtins first.

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=sorting.go) -->
<!-- The below code snippet is automatically added from sorting.go -->
```go
package main

import (
	"fmt"
	"sort"
)

func main() {

	// Sort methods are specific to the builtin type;
	// here's an example for strings
	// Note that sorting is in-place,
	// so it changes the given slice and doesn't return a new one
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	// An example of sorting ints
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	// We can also use sort to check if a slice is already in sorted order
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}
```
<!-- AUTO-GENERATED-CONTENT:END -->

- Running our program prints the sorted string and int slices and true as the result of our AreSorted test.

```bash
$ go run sorting.go
```


## Sorting by Functions

> Sometimes we'll want to sort a collection by something other than its natural order.
> For example, suppose we wanted to sort strings by their length instead of alphabetically.
> Here's an example of custom sorts in Go.

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=sorting-by-functions.go) -->
<!-- The below code snippet is automatically added from sorting-by-functions.go -->
```go
package main

import (
	"fmt"
	"sort"
)

// In order to sort by a custom function in Go,
// we need a corresponding type.
// Here we've created a byLength type that is just an alias for the builtin []string type
type byLength []string

// We implement sort.Interface - Len, Less, and Swap - on our type so we can use the sort package's generic Sort function
// Len and Swap will usually be similar across types and Less will hold the actual custom sorting logic
// In our case we want to sort in order of increasing string length,
// so we use len(s[i]) and len(s[j]) here
func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

// With all of this in place,
// we can now implement our custom sort by converting the original fruits slice to byLength,
// and then use sort.Sort on that typed slice
func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
```
<!-- AUTO-GENERATED-CONTENT:END -->

- Running our program shows a list sorted by string length, as desired.

```bash
$ go run sorting-by-functions.go
```

- By following this same pattern of creating a custom type, implementing the three Interface methods on that type, and then calling sort.Sort on a collection of that custom type, we can sort Go slices by arbitrary functions.


## Panic

> A panic typically means something went unexpectedly wrong.
> Mostly we use it to fail fast on errors that shouldn't occur during normal operation,
> or that we aren't prepared to handle gracefully

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=panic.go) -->
<!-- The below code snippet is automatically added from panic.go -->
```go
package main

import "os"

func main() {

	// We'll use panic throughout this site to check for unexpected errors
	// This is the only program on the site designed to panic
	panic("a problem")

	// A common use of panic is to abort if a function returns an error value that we don't know how to (or want to) handle
	// Here's an example of panicking if we get an unexpected error when creating a new file
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
```
<!-- AUTO-GENERATED-CONTENT:END -->

- Running this program will cause it to panic, print an error message and goroutine traces, and exit with a non-zero status.

```bash
$ go run panic.go
```

- Note that unlike some languages which use exceptions for handling of many errors, in Go it is idiomatic to use error-indicating return values wherever possible.


## Defer

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=defer.go) -->
<!-- The below code snippet is automatically added from defer.go -->
```go
package main

import (
	"fmt"
	"os"
)

func main() {

	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")

}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
```
<!-- AUTO-GENERATED-CONTENT:END -->

```bash
$ go run defer.go
```


## References

- [Go by Example](https://gobyexample.com/)
