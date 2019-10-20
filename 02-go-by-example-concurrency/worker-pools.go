package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Here is the worker, of which we will run several concurrent instances
// These workers will receive work on the jobs channel and send the corresponding results on results
// We will sleep randomly per job to simulate an expensive task
func worker(id int, jobs <-chan int, results chan<- int) {

	for job := range jobs {
		fmt.Println("Worker :", id, " -> Starting Job :", job)

		duration := time.Duration(rand.Intn(3000)) * time.Millisecond
		time.Sleep(duration)
		fmt.Println("Worker :", id, " -> Finished Job :", job, " -> Duration :", duration)

		results <- job * 2
	}
}

func main() {

	// In order to use our pool of workers we need to send them work and collect their results
	// We make 2 channels for this
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// This starts up 3 workers, initially blocked because there are no jobs yet
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Here we send 5 jobs and then close that channel to indicate that's all the work we have
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// Finally we collect all the results of the work
	// This also ensures that the worker goroutines have finished
	// An alternative way to wait for multiple goroutines is to use a WaitGroup
	for a := 1; a <= 5; a++ {
		<-results
	}
}
