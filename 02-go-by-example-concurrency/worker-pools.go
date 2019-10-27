package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Here is the "worker", of which we will run several concurrent instances
// These will receive "jobs" and send the corresponding "results"
func worker(id int, jobs <-chan int, results chan<- int) {

	for job := range jobs {
		fmt.Println("Worker :", id, " -> Starting Job :", job)

		// Sleep randomly per job to simulate an expensive task
		duration := time.Duration(rand.Intn(3000)) * time.Millisecond
		time.Sleep(duration)
		fmt.Println("Worker :", id, " -> Finished Job :", job, " -> Duration :", duration)

		results <- job
	}
}

func main() {

	// In order to use our pool of workers,
	// we need to send them "jobs" and collect their "results"
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// Start up "3" workers,
	// initially blocked because there are no jobs yet
	for id := 1; id <= 3; id++ {
		go worker(id, jobs, results)
	}

	// Here we send "5" jobs,
	// and then "close" that channel to indicate that is all the work we have
	for job := 1; job <= 5; job++ {
		jobs <- job
	}
	close(jobs)

	// Finally, we collect all the "results" of the work
	// This also ensures that the worker goroutines have finished
	for count := 1; count <= 5; count++ {
		fmt.Println("Received Result :", <-results)
	}
}
