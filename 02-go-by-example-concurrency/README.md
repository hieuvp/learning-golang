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
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// Suppose we have a function call f(s)
	// Here's how we’d call that in the usual way, running it synchronously
	f("direct")

	// To invoke this function in a goroutine, use go f(s)
	// This new goroutine will execute concurrently with the calling one
	go f("goroutine")

	// You can also start a goroutine for an anonymous function call
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// Our two function calls are running asynchronously in separate goroutines now
	// Wait for them to finish (for a more robust approach, use a WaitGroup)
	time.Sleep(time.Second)
	fmt.Println("done")
}
```
<!-- AUTO-GENERATED-CONTENT:END -->

- When we run this program, we see the output of the blocking call first, then the interleaved output of the two goroutines. This interleaving reflects the goroutines being run concurrently by the Go runtime.

```bash
$ go run goroutines.go

# direct : 0
# direct : 1
# direct : 2
# goroutine : 0
# goroutine : 1
# goroutine : 2
# going
# done
```


## Channels

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=channels.go) -->
<!-- The below code snippet is automatically added from channels.go -->
```go
package main

import "fmt"

func main() {

	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)
}
```
<!-- AUTO-GENERATED-CONTENT:END -->

```bash
$ go run channels.go

# ping
```


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
