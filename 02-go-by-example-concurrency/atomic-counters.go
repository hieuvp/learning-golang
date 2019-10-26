package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	// Use an unsigned integer to represent our always-positive number
	var counter uint64
	var atomicCounter uint64

	// A "WaitGroup" will help us wait for all goroutines to finish their work
	var wg sync.WaitGroup

	// We will start "50" goroutines that
	// each increment the counter exactly "1000" times
	for i := 0; i < 50; i++ {
		wg.Add(1)

		// To atomically increment the counter we use "AddUint64",
		// giving it the memory address of our "atomicCounter" counter with the "&" syntax
		go func() {
			for j := 0; j < 1000; j++ {
				counter++
				atomic.AddUint64(&atomicCounter, 1)
			}
			wg.Done()
		}()
	}

	// Wait until all the goroutines are done
	wg.Wait()

	// It is safe to access "counter" and "atomicCounter" now
	// because we know no other goroutine is writing to it
	// Reading atomics safely while they are being updated is also possible,
	// using functions like "atomic.LoadUint64"
	fmt.Println("counter       =", counter)
	fmt.Println("atomicCounter =", atomicCounter)
}
