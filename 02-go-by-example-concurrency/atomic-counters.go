package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	// We will use an unsigned integer to represent our (always-positive) counter
	var ops uint64
	var atomicOps uint64

	// A "WaitGroup" will help us wait for all goroutines to finish their work
	var wg sync.WaitGroup

	// We will start "50" goroutines that each increment the counter exactly 1000 times
	for i := 0; i < 50; i++ {
		wg.Add(1)

		// To atomically increment the counter we use "AddUint64",
		// giving it the memory address of our "atomicOps" counter with the "&" syntax
		go func() {
			for counter := 0; counter < 1000; counter++ {
				ops += 1
				atomic.AddUint64(&atomicOps, 1)
			}
			wg.Done()
		}()
	}

	// Wait until all the goroutines are done
	wg.Wait()

	// It is safe to access "atomicOps" now because we know no other goroutine is writing to it
	// Reading atomics safely while they are being updated is also possible,
	// using functions like "atomic.LoadUint64"
	fmt.Println("ops =", ops)
	fmt.Println("atomicOps =", atomicOps)
}
