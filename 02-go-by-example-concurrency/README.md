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

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=select.go) -->
<!-- The below code snippet is automatically added from select.go -->
```go
package main

import (
	"fmt"
	"time"
)

func main() {

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
```
<!-- AUTO-GENERATED-CONTENT:END -->

```bash
$ go run select.go

# received one
# received two
```


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

	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
```
<!-- AUTO-GENERATED-CONTENT:END -->

```bash
$ go run timeouts.go

# timeout 1
# result 2
```


## Non-Blocking Channel Operations

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=non-blocking-channel-operations.go) -->
<!-- The below code snippet is automatically added from non-blocking-channel-operations.go -->
```go
package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
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
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done
}
```
<!-- AUTO-GENERATED-CONTENT:END -->

```bash
$ go run closing-channels.go

# sent job 1
# sent job 2
# sent job 3
# sent all jobs
# received job 1
# received job 2
# received job 3
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
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
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

## Stateful Goroutines

## Sorting

## Sorting by Functions

## Panic

## Defer


## References

- [Go by Example](https://gobyexample.com/)
