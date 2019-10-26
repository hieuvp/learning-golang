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
