package main

import "fmt"

func main() {

	// We will use a "jobs" channel to communicate work to be done
	// from the main() goroutine to a worker goroutine
	// When we have no more jobs for the worker we will close the "jobs" channel
	jobs := make(chan int, 5)
	done := make(chan bool)

	// Here is the worker goroutine
	go func() {

		// Repeatedly receives from jobs
		for {

			// If jobs has been closed and all values in the channel have already been received,
			// The "more" value will be "false"
			// We use this to notify on done when we have worked all our jobs
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// This sends 3 jobs to the worker over the "jobs" channel, then closes it
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}

	// "all goroutines are asleep - deadlock!" when omitted
	close(jobs)
	fmt.Println("sent all jobs")

	// We await the worker using the synchronization approach we saw earlier
	<-done
}
