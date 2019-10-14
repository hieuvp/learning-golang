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

	// Suppose we have a function call "f(s)"
	// Here's how we'd call that in the usual way, running it "synchronously"
	f("Direct")

	// To invoke this function in a goroutine, use "go f(s)"
	// This new goroutine will execute concurrently with the calling one
	go f("Goroutine")

	// You can also start a goroutine for an anonymous function call
	go func(msg string) {
		fmt.Println(msg)
	}("Going")

	// Our two function calls above are running "asynchronously" in separate goroutines now

	// Wait for them to finish
	// "time.Sleep" pauses the current goroutine for at least the "Duration"
	fmt.Println("Sleep with Duration:", time.Second)
	time.Sleep(time.Second)

	fmt.Println("Done")
}
```
<!-- AUTO-GENERATED-CONTENT:END -->

- When we run this program, we see the output of the blocking call first, then the interleaved output of the two goroutines. This interleaving reflects the goroutines being run concurrently by the Go runtime.

```bash
$ go run goroutines.go

# Direct : 0
# Direct : 1
# Direct : 2
# Direct : 3
# Goroutine : 0
# Sleep with Duration: 1s
# Goroutine : 1
# Goroutine : 2
# Goroutine : 3
# Going
# Done
```


## Channels

> **Channels** are the pipes that connect concurrent goroutines.
> You can send values into channels from one goroutine and receive those values into another goroutine.

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=channels.go) -->
<!-- The below code snippet is automatically added from channels.go -->
```go
package main

import "fmt"

func main() {

	// Create a new channel with "make(chan value-type)"
	// Channels are typed by the values they convey
	messages := make(chan string)

	// Send a value into a channel using the channel "<-" syntax
	// Here we send "ping" to the messages channel we made above, from a new goroutine
	go func() { messages <- "ping" }()

	// The "<-channel" syntax receives a value from the channel
	// Here we will receive the "ping" message we sent above and print it out
	msg := <-messages
	fmt.Println(msg)
}
```
<!-- AUTO-GENERATED-CONTENT:END -->

- When we run the program the "ping" message is successfully passed from one goroutine to another via our channel.

```bash
$ go run channels.go

# ping
```

- By default sends and receives block until both the sender and receiver are ready. This property allowed us to wait at the end of our program for the "ping" message without having to use any other synchronization.


## Channel Buffering

## Channel Synchronization

## Channel Directions

## Select

## Timeouts

## Non-Blocking Channel Operations

## Closing Channels

## Range over Channels

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
