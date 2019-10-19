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

			j, more := <-jobs

			// If "jobs" has been closed and all values in the channel have already been received,
			// the "more" value will be "false"
			if more {
				fmt.Println("received job :", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// Send 3 jobs to the "worker" over the "jobs" channel
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job :", j)
	}

	// If a channel is closed, no values can be sent on it
	// "close" a channel to indicate "completion"
	close(jobs)
	fmt.Println("sent all jobs")

	// We await the "worker" using the synchronization approach
	<-done
}
