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

- **Sends** to a buffered channel block only when the buffer is **full**.
- **Receives** block when the buffer is **empty**.


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

	// In contrast,
	// "make(chan string)" as equivalent to "make(chan string, 0)" is an "unbuffered channel"
	// The sender will block on the channel until the receiver receives the data from the channel

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

> `Timer` is for when you want to do something once in the future.

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=timers.go) -->
<!-- The below code snippet is automatically added from timers.go -->
```go
package main

import (
	"fmt"
	"time"
)

func main() {

	// "Timer" represents a single event in the future
	// You tell the "Timer" how long you want to wait
	timer1 := time.NewTimer(2 * time.Second)

	// It provides a channel that will send a value indicating when the "Timer" expired
	<-timer1.C
	fmt.Println("Timer 1 expired")

	// If you just wanted to wait, you could have used "time.Sleep"
	// One reason a "Timer" maybe useful is that,
	// you can "stop" the "Timer" before it expires
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

> `Ticker` is for when you want to do something repeatedly at regular intervals.

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=tickers.go) -->
<!-- The below code snippet is automatically added from tickers.go -->
```go
package main

import (
	"fmt"
	"time"
)

func main() {

	// "Ticker" uses a similar mechanism to "Timer": a channel that is sent values
	// Here we await the values as they arrive "every 500ms"
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at :", t)
			}
		}
	}()

	// "Ticker" can be stopped like "Timer"
	// Once a "Ticker" is stopped, it won't receive any more values on its channel
	time.Sleep(1600 * time.Millisecond)

	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")

	time.Sleep(2000 * time.Millisecond)
}
```
<!-- AUTO-GENERATED-CONTENT:END -->

```bash
$ go run tickers.go

# Tick at : 2019-10-23 12:05:26.605803 +0700 +07 m=+0.504344617
# Tick at : 2019-10-23 12:05:27.106393 +0700 +07 m=+1.004949855
# Tick at : 2019-10-23 12:05:27.60466 +0700 +07 m=+1.503231861

# Ticker stopped
```


## Worker Pools

> In this example, we will look at how to implement a **worker pool** using **goroutines** and **channels**.

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=worker-pools.go) -->
<!-- The below code snippet is automatically added from worker-pools.go -->
```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Here is the "worker", of which we will run several concurrent instances
// These will receive "jobs" and send the corresponding "results"
func worker(id int, jobs <-chan int, results chan<- int) {

	for job := range jobs {
		fmt.Println("Worker :", id, " -> Starting Job :", job)

		// Sleep randomly per job to simulate an expensive task
		duration := time.Duration(rand.Intn(3000)) * time.Millisecond
		time.Sleep(duration)
		fmt.Println("Worker :", id, " -> Finished Job :", job, " -> Duration :", duration)

		results <- job
	}
}

func main() {

	// In order to use our pool of workers,
	// we need to send them "jobs" and collect their "results"
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// Start up "3" workers,
	// initially blocked because there are no jobs yet
	for id := 1; id <= 3; id++ {
		go worker(id, jobs, results)
	}

	// Here we send "5" jobs,
	// and then "close" that channel to indicate that is all the work we have
	for job := 1; job <= 5; job++ {
		jobs <- job
	}
	close(jobs)

	// Finally, we collect all the "results" of the work
	// This also ensures that the worker goroutines have finished
	for count := 1; count <= 5; count++ {
		fmt.Println("Received Result :", <-results)
	}
}
```
<!-- AUTO-GENERATED-CONTENT:END -->

```bash
$ time go run worker-pools.go

# Worker : 3  -> Starting Job : 1
# Worker : 2  -> Starting Job : 3
# Worker : 1  -> Starting Job : 2
# Worker : 1  -> Finished Job : 2  -> Duration : 1.847s
# Worker : 1  -> Starting Job : 4
# Received Result : 2
# Worker : 2  -> Finished Job : 3  -> Duration : 1.887s
# Worker : 2  -> Starting Job : 5
# Received Result : 3
# Worker : 1  -> Finished Job : 4  -> Duration : 59ms
# Received Result : 4
# Worker : 3  -> Finished Job : 1  -> Duration : 2.081s
# Received Result : 1
# Worker : 2  -> Finished Job : 5  -> Duration : 1.081s
# Received Result : 5

# go run worker-pools.go  0.28s user 0.21s system 14% cpu 3.264 total
```


## WaitGroups

> To wait for multiple goroutines to finish, we can use a `WaitGroup`.

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=wait-groups.go) -->
<!-- The below code snippet is automatically added from wait-groups.go -->
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

// Because we are going to use a method "func (wg *WaitGroup) Done()",
// so "WaitGroup" must be passed by "pointer"
func worker(id int, wg *sync.WaitGroup) {
	fmt.Printf("Worker : %d -> Starting\n", id)

	// Sleep to simulate an expensive task
	time.Sleep(time.Second)
	fmt.Printf("Worker : %d -> Done\n", id)

	// Notify the "WaitGroup" that this "worker" is "done"
	wg.Done()
}

func main() {

	// This "WaitGroup" is used to wait for all the goroutines launched here to finish
	var wg sync.WaitGroup

	// Launch several goroutines and increase the "WaitGroup" "counter" for each
	for id := 1; id <= 5; id++ {
		wg.Add(1)
		go worker(id, &wg)
	}

	// Block until the "WaitGroup" "counter" goes back to "0",
	// meaning all the workers notified that they are "done"
	wg.Wait()
}
```
<!-- AUTO-GENERATED-CONTENT:END -->

```bash
$ go run wait-groups.go

# Worker : 5 -> Starting
# Worker : 2 -> Starting
# Worker : 1 -> Starting
# Worker : 4 -> Starting
# Worker : 3 -> Starting

# Worker : 5 -> Done
# Worker : 1 -> Done
# Worker : 3 -> Done
# Worker : 4 -> Done
# Worker : 2 -> Done
```


## Rate Limiting

- **Rate limiting** is an important mechanism for controlling resource utilization and maintaining quality of service.
- Go elegantly supports **rate limiting** with **goroutines**, **channels**, and **tickers**.

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=rate-limiting.go) -->
<!-- The below code snippet is automatically added from rate-limiting.go -->
```go
package main

import (
	"fmt"
	"strings"
	"time"
)

func getMonotonicClock() string {
	now := time.Now().String()
	separator := "m=+"
	return strings.Split(now, separator)[1]
}

func main() {

	// Suppose we want to limit our handling of incoming requests
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// The "limiter" is the regulator in our "rate limiting" scheme
	duration := 200 * time.Millisecond
	fmt.Println("Initializing Ticker for every", duration)
	limiter := time.Tick(duration)

	// By blocking on a receive from the "limiter" channel before serving each request,
	// we limit ourselves to "1" request every "200ms"
	for request := range requests {
		<-limiter
		fmt.Println("Request", request, ": at", getMonotonicClock())
	}
	fmt.Println()

	// We may want to allow short bursts of requests in our "rate limiting" scheme
	// while preserving the overall rate limit

	// This "burstyLimiter" channel will allow bursts of up to "3" events
	burstyLimiter := make(chan time.Time, 3)
	for i := 1; i <= 3; i++ {
		burstyLimiter <- time.Now()
	}
	fmt.Printf("Initial Limiter length     : %d\n\n", len(burstyLimiter))

	// Every "200ms" we will try to add a new value to "burstyLimiter",
	// up to its limit of "3"
	go func() {
		fmt.Printf("Initializing Ticker for every %s\n\n", duration)

		for t := range time.Tick(duration) {
			fmt.Println("Ticking                    : at", getMonotonicClock())
			fmt.Println("Increasing Limiter length  :", len(burstyLimiter), "+ 1")

			// Blocked when "burstyLimiter" reaches its limited buffered capacity
			burstyLimiter <- t
			fmt.Println("Increased Limiter length   : at", getMonotonicClock())
		}
	}()

	// Now simulating "15" incoming requests
	burstyRequests := make(chan int, 15)
	for i := 1; i <= 15; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	// The first "3" requests will benefit from the burst capability of "burstyLimiter"
	for request := range burstyRequests {

		if duration := 5 * time.Second; request == 7 {
			fmt.Printf("Go to Sleep for %s...\n\n", duration)
			time.Sleep(duration)
			fmt.Printf("\nLimiter length after Sleep : %d\n\n", len(burstyLimiter))
		}

		<-burstyLimiter
		fmt.Println("Decreased Limiter length   :", len(burstyLimiter))
		fmt.Printf("Request %2d                 : at %s\n\n", request, getMonotonicClock())
	}
}
```
<!-- AUTO-GENERATED-CONTENT:END -->

```bash
$ go run rate-limiting.go

# Initializing Ticker for every 200ms
# Request 1 : at 0.203333473
# Request 2 : at 0.405552978
# Request 3 : at 0.604721689
# Request 4 : at 0.800712096
# Request 5 : at 1.002624976

# Initial Limiter length     : 3

# Decreased Limiter length   : 2
# Request  1                 : at 1.002771540

# Decreased Limiter length   : 1
# Request  2                 : at 1.002788950

# Decreased Limiter length   : 0
# Request  3                 : at 1.002797462

# Initializing Ticker for every 200ms

# Ticking                    : at 1.204327299
# Increasing Limiter length  : 0 + 1
# Increased Limiter length   : at 1.204377655
# Decreased Limiter length   : 0
# Request  4                 : at 1.204432530

# Ticking                    : at 1.407059592
# Increasing Limiter length  : 0 + 1
# Decreased Limiter length   : 0
# Increased Limiter length   : at 1.407129247
# Request  5                 : at 1.407216867

# Ticking                    : at 1.607240448
# Increasing Limiter length  : 0 + 1
# Increased Limiter length   : at 1.607272617
# Decreased Limiter length   : 0
# Request  6                 : at 1.607314612

# Go to Sleep for 5s...

# Ticking                    : at 1.802901295
# Increasing Limiter length  : 0 + 1
# Increased Limiter length   : at 1.802959804
# Ticking                    : at 2.004420485
# Increasing Limiter length  : 1 + 1
# Increased Limiter length   : at 2.004480205
# Ticking                    : at 2.203566231
# Increasing Limiter length  : 2 + 1
# Increased Limiter length   : at 2.203660086
# Ticking                    : at 2.406305773
# Increasing Limiter length  : 3 + 1

# Limiter length after Sleep : 3

# Decreased Limiter length   : 3
# Request  7                 : at 6.608070576

# Decreased Limiter length   : 2
# Request  8                 : at 6.608093022

# Decreased Limiter length   : 1
# Increased Limiter length   : at 6.608083460
# Ticking                    : at 6.608112758
# Increasing Limiter length  : 1 + 1
# Request  9                 : at 6.608102555

# Decreased Limiter length   : 1
# Increased Limiter length   : at 6.608122113
# Request 10                 : at 6.608159171

# Decreased Limiter length   : 0
# Request 11                 : at 6.608238821

# Ticking                    : at 6.805545798
# Increasing Limiter length  : 0 + 1
# Increased Limiter length   : at 6.805597046
# Decreased Limiter length   : 0
# Request 12                 : at 6.805628023

# Ticking                    : at 7.005515089
# Increasing Limiter length  : 0 + 1
# Increased Limiter length   : at 7.005568954
# Decreased Limiter length   : 0
# Request 13                 : at 7.005641972

# Ticking                    : at 7.208208684
# Increasing Limiter length  : 0 + 1
# Increased Limiter length   : at 7.208269061
# Decreased Limiter length   : 0
# Request 14                 : at 7.208313312

# Ticking                    : at 7.403927743
# Increasing Limiter length  : 0 + 1
# Increased Limiter length   : at 7.403978700
# Decreased Limiter length   : 0
# Request 15                 : at 7.404023068
```


## Atomic Counters

- The primary mechanism for **managing state** in Go is **communication over channels**, we saw this for example with [**Worker Pools**](#worker-pools).
- There are a few other options for **managing state** though, here we will look at using the `sync/atomic` package for **atomic counters** accessed by multiple goroutines.

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=atomic-counters.go) -->
<!-- The below code snippet is automatically added from atomic-counters.go -->
```go
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	// Use unsigned integers to represent our always-positive numbers
	var nonAtomicCounter uint64
	var atomicCounter uint64

	// A "WaitGroup" will help us wait for all goroutines to finish their work
	var wg sync.WaitGroup

	// Start "50" goroutines that each increases the counters exactly "1000" times
	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for j := 0; j < 1000; j++ {
				nonAtomicCounter++

				// Atomically increment
				atomic.AddUint64(&atomicCounter, 1)
			}

			wg.Done()
		}()
	}

	// Wait until all the goroutines are done
	wg.Wait()

	// It is safe to access our counters now
	// because we know no other goroutine is writing to them
	fmt.Println("nonAtomicCounter :", nonAtomicCounter)
	fmt.Println("atomicCounter    :", atomicCounter)

	// Reading atomics safely while they are being updated is also possible
	fmt.Println("atomic.LoadUint64(&atomicCounter) =", atomic.LoadUint64(&atomicCounter))

}
```
<!-- AUTO-GENERATED-CONTENT:END -->

```bash
$ go run atomic-counters.go

# nonAtomicCounter : 17806
# atomicCounter    : 50000

# atomic.LoadUint64(&atomicCounter) = 50000
```

- With `nonAtomicCounter`, we would likely get a different number, changing between runs, because the goroutines interfere with each other, **race condition**.


## Mutexes


## Stateful Goroutines


## Sorting


## Sorting by Functions


## Panic


## Defer


## References

- [Go by Example](https://gobyexample.com/)
