package main

import (
	"fmt"
	"strings"
	"time"
)

func getMonotonicClock() string {
	now := time.Now().String()
	separator := "m=+"
	return strings.Split(now, separator)[1]
}

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
		fmt.Println("Request", request, ": at", getMonotonicClock())
	}
	fmt.Println("\n")

	// We may want to allow short bursts of requests in our rate limiting scheme
	// while preserving the overall rate limit
	// This "burstyLimiter" channel will allow bursts of up to "3" events
	burstyLimiter := make(chan time.Time, 3)

	// Fill up the channel to represent allowed bursting
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}
	fmt.Println("Initial Limiter capacity     :", len(burstyLimiter))

	// Every "200ms" we will try to add a new value to "burstyLimiter",
	// up to its limit of "3"
	go func() {
		duration := 200 * time.Millisecond
		fmt.Printf("\nInitializing Ticker for every %s\n\n", duration)

		for t := range time.Tick(duration) {
			fmt.Println("Ticking                      : at", getMonotonicClock())
			fmt.Println("Increasing Limiter capacity  :", len(burstyLimiter), "+ 1")

			// Blocking when "burstyLimiter" reaches its maximum buffered capacity
			burstyLimiter <- t
			fmt.Println("Increased Limiter capacity   : at", getMonotonicClock())
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
		fmt.Println("Decreased Limiter capacity   :", len(burstyLimiter))

		if duration := 5 * time.Second; request == 7 {
			fmt.Printf("\nGo to Sleep for %s...\n\n", duration)
			time.Sleep(duration)
			fmt.Printf("\nLimiter capacity after Sleep : %d\n\n", len(burstyLimiter))
		}

		fmt.Printf("Request %2d                   : at %s\n", request, getMonotonicClock())
	}
}
