package main

import (
	"fmt"
	"time"
)

func main() {

	// Basic rate limiting
	// Suppose we want to limit our handling of incoming requests
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// This "limiter" is the regulator in our rate limiting scheme
	limiter := time.Tick(200 * time.Millisecond)

	// By blocking on a receive from the "limiter" channel before serving each request,
	// we limit ourselves to "1" request every "200ms"
	for request := range requests {
		<-limiter
		fmt.Println("Request", request, ":", time.Now())
	}
	fmt.Println()

	// We may want to allow short bursts of requests in our rate limiting scheme
	// while preserving the overall rate limit
	// This "burstyLimiter" channel will allow bursts of up to "3" events
	burstyLimiter := make(chan time.Time, 3)

	// Fill up the channel to represent allowed bursting
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}
	fmt.Println("Initial burstyLimiter capacity :", len(burstyLimiter))

	// Every "200ms" we will try to add a new value to "burstyLimiter",
	// up to its limit of "3"
	go func() {
		duration := 200 * time.Millisecond
		fmt.Printf("Initializing ticks in every %s\n", duration)

		for t := range time.Tick(duration) {
			fmt.Println("Tick :", time.Now())
			fmt.Println("burstyLimiter capacity will be increased by 1 from", len(burstyLimiter))

			// The send will be blocked when reaches maximum buffered length,
			// in this case "3"
			burstyLimiter <- t
			fmt.Println("burstyLimiter capacity increased :", time.Now())

		}
	}()

	// Now simulate "15" more incoming requests
	// The first "3" of these will benefit from the burst capability of "burstyLimiter"
	burstyRequests := make(chan int, 15)
	for i := 1; i <= 15; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for request := range burstyRequests {
		<-burstyLimiter

		if duration := 5 * time.Second; request == 7 {
			fmt.Printf("Sleeping for %s...\n", duration)
			time.Sleep(duration)
			fmt.Println("burstyLimiter capacity after sleeping :", len(burstyLimiter))
		}

		fmt.Printf("Request %2d : %s\n", request, time.Now())
	}
}
