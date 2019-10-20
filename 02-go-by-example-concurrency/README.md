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

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=timers.go) -->
<!-- The below code snippet is automatically added from timers.go -->
```go
package main

import (
	"fmt"
	"time"
)

func main() {

	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	fmt.Println("Timer 1 expired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}
```
<!-- AUTO-GENERATED-CONTENT:END -->

```bash
$ go run timers.go

# Timer 1 expired
# Timer 2 stopped
```


## Tickers

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=tickers.go) -->
<!-- The below code snippet is automatically added from tickers.go -->
```go
package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
```
<!-- AUTO-GENERATED-CONTENT:END -->

```bash
$ go run tickers.go

# Tick at 2012-09-23 11:29:56.487625 -0700 PDT
# Tick at 2012-09-23 11:29:56.988063 -0700 PDT
# Tick at 2012-09-23 11:29:57.488076 -0700 PDT
# Ticker stopped
```


## Worker Pools

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=worker-pools.go) -->
<!-- The below code snippet is automatically added from worker-pools.go -->
```go
package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= 5; a++ {
		<-results
	}
}
```
<!-- AUTO-GENERATED-CONTENT:END -->

```bash
$ time go run worker-pools.go

# worker 3 started  job 1
# worker 1 started  job 2
# worker 2 started  job 3
# worker 1 finished job 2
# worker 2 finished job 3
# worker 1 started  job 4
# worker 3 finished job 1
# worker 2 started  job 5
# worker 2 finished job 5
# worker 1 finished job 4

# go run worker-pools.go  0.26s user 0.21s system 20% cpu 2.283 total
```


## WaitGroups

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=wait-groups.go) -->
<!-- The below code snippet is automatically added from wait-groups.go -->
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)

	wg.Done()
}

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
}
```
<!-- AUTO-GENERATED-CONTENT:END -->

```bash
$ go run wait-groups.go

# Worker 5 starting
# Worker 2 starting
# Worker 3 starting
# Worker 4 starting
# Worker 1 starting
# Worker 1 done
# Worker 2 done
# Worker 3 done
# Worker 4 done
# Worker 5 done
```


## Rate Limiting

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=rate-limiting.go) -->
<!-- The below code snippet is automatically added from rate-limiting.go -->
```go
package main

import (
	"fmt"
	"time"
)

func main() {

	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
```
<!-- AUTO-GENERATED-CONTENT:END -->

```bash
$ go run rate-limiting.go

# request 1 2012-10-19 00:38:18.687438 +0000 UTC
# request 2 2012-10-19 00:38:18.887471 +0000 UTC
# request 3 2012-10-19 00:38:19.087238 +0000 UTC
# request 4 2012-10-19 00:38:19.287338 +0000 UTC
# request 5 2012-10-19 00:38:19.487331 +0000 UTC

# request 1 2012-10-19 00:38:20.487578 +0000 UTC
# request 2 2012-10-19 00:38:20.487645 +0000 UTC
# request 3 2012-10-19 00:38:20.487676 +0000 UTC
# request 4 2012-10-19 00:38:20.687483 +0000 UTC
# request 5 2012-10-19 00:38:20.887542 +0000 UTC
```


## Atomic Counters


## Mutexes


## Stateful Goroutines


## Sorting


## Sorting by Functions


## Panic


## Defer


## References

- [Go by Example](https://gobyexample.com/)
