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
