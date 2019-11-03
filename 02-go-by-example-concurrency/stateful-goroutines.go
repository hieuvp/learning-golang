package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// Our "state" will be owned by "a single goroutine"
// This will guarantee that the data is never corrupted with concurrent access

// In order to read or write that "state",
// other goroutines will send messages to the owning goroutine
// and receive corresponding replies

// These "readOperation" and "writeOperation" structs
// encapsulate those requests and a way for the owning goroutine to respond
type readOperation struct {
	key      int
	response chan int
}
type writeOperation struct {
	key      int
	value    int
	response chan bool
}

func main() {

	// As before, we will count how many operations we perform
	var readOps uint64
	var writeOps uint64

	// The "reads" and "writes" channels will be used by other goroutines
	// to issue read and write requests, respectively
	reads := make(chan readOperation)
	writes := make(chan writeOperation)

	// Here is the goroutine that owns the "state",
	// which is a "map" as in the previous example but now private to this "stateful goroutine"
	go func() {
		var state = make(map[int]int)

		// Repeatedly "select" on the "reads" and "writes" channels,
		// responding to requests as they arrive
		for {

			select {
			case read := <-reads:
				// Send desired value
				read.response <- state[read.key]
			case write := <-writes:
				state[write.key] = write.value
				// Indicate success
				write.response <- true
			}
		}
	}()

	// This starts "100" goroutines to
	// issue reads to the state-owning goroutine via the "reads" channel
	for r := 0; r < 100; r++ {
		go func() {
			for {

				// Each read requires constructing a "readOperation"
				read := readOperation{
					key:      rand.Intn(5),
					response: make(chan int)}

				// Sending it over the "reads" channel
				reads <- read

				// Receiving the result over the provided "read.response" channel
				<-read.response

				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// We start "10" writes as well, using a similar approach
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeOperation{
					key:      rand.Intn(5),
					value:    rand.Intn(100),
					response: make(chan bool)}

				writes <- write
				<-write.response

				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Let the goroutines work for "3s"
	time.Sleep(3 * time.Second)

	// Finally, capture and report the operation counts
	finalReadOps := atomic.LoadUint64(&readOps)
	finalWriteOps := atomic.LoadUint64(&writeOps)
	fmt.Println("readOps  :", finalReadOps)
	fmt.Println("writeOps :", finalWriteOps)
}
