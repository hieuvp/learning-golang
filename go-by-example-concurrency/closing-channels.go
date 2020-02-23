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
			job, more := <-jobs

			// The "more" value will be "false",
			// if "jobs" has been closed and
			// all values in the channel have already been received
			if more {
				fmt.Println("received job :", job)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// Send "3" jobs to the "worker" over the "jobs" channel
	for job := 1; job <= 3; job++ {
		jobs <- job
		fmt.Println("sent job :", job)
	}

	// If a channel is closed, no values can be sent on it
	// "close" a channel to indicate "completion"
	close(jobs)
	fmt.Println("sent all jobs")

	// We await the "worker" using the synchronization approach
	<-done
}
