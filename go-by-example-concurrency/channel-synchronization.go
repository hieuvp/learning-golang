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
