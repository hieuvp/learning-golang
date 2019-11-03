package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	// Previously, we saw how to manage "simple" counter state using "atomic" operations
	// For more "complex" state (e.g. map),
	// we can use a "Mutex" to safely access data across multiple goroutines
	var state = make(map[int]int)

	// A "Mutex" is used to provide a locking mechanism to ensure that
	// only one goroutine is running the critical section of code at any point of time
	// to prevent race condition from happening

	// This "mutex" will synchronize access to "state"
	var mutex = &sync.Mutex{}

	// We will keep track of how many read and write operations we do
	var readOps uint64
	var writeOps uint64

	// Here we start "100" goroutines to execute repeated reads against the "state"
	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {

				// For each read, we pick a random "key" to access
				key := rand.Intn(5)

				// "Lock()" the "mutex" to ensure exclusive access to the "state"
				mutex.Lock()

				// Read the value at the chosen "key"
				total += state[key]

				// "Unlock()" the "mutex"
				mutex.Unlock()

				// Increase the "readOps" count
				atomic.AddUint64(&readOps, 1)

				// Wait a bit between reads
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// We will also start "10" goroutines to simulate writes
	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				value := rand.Intn(100)

				mutex.Lock()
				state[key] = value
				mutex.Unlock()

				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Let our goroutines work on the "state" and "mutex" for "3s"
	time.Sleep(3 * time.Second)

	// Take and report final operation counts
	finalReadOps := atomic.LoadUint64(&readOps)
	finalWriteOps := atomic.LoadUint64(&writeOps)
	fmt.Println("readOps  :", finalReadOps)
	fmt.Println("writeOps :", finalWriteOps)

	// Show how the final "state" ended up
	mutex.Lock()
	fmt.Println("state    :", state)
	mutex.Unlock()
}
