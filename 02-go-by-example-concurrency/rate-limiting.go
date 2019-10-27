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

	// Suppose we want to limit our handling of incoming requests
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// The "limiter" is the regulator in our "rate limiting" scheme
	duration := 200 * time.Millisecond
	fmt.Println("Initializing Ticker for every", duration)
	limiter := time.Tick(duration)

	// By blocking on a receive from the "limiter" channel before serving each request,
	// we limit ourselves to "1" request every "200ms"
	for request := range requests {
		<-limiter
		fmt.Println("Request", request, ": at", getMonotonicClock())
	}
	fmt.Println()

	// We may want to allow short bursts of requests in our "rate limiting" scheme
	// while preserving the overall rate limit

	// This "burstyLimiter" channel will allow bursts of up to "3" events
	burstyLimiter := make(chan time.Time, 3)
	for i := 1; i <= 3; i++ {
		burstyLimiter <- time.Now()
	}
	fmt.Printf("Initial Limiter length     : %d\n\n", len(burstyLimiter))

	// Every "200ms" we will try to add a new value to "burstyLimiter",
	// up to its limit of "3"
	go func() {
		fmt.Printf("Initializing Ticker for every %s\n\n", duration)

		for t := range time.Tick(duration) {
			fmt.Println("Ticking                    : at", getMonotonicClock())
			fmt.Println("Increasing Limiter length  :", len(burstyLimiter), "+ 1")

			// Blocked when "burstyLimiter" reaches its limited buffered capacity
			burstyLimiter <- t
			fmt.Println("Increased Limiter length   : at", getMonotonicClock())
		}
	}()

	// Now simulating "15" incoming requests
	burstyRequests := make(chan int, 15)
	for i := 1; i <= 15; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	// The first "3" requests will benefit from the burst capability of "burstyLimiter"
	for request := range burstyRequests {

		if duration := 5 * time.Second; request == 7 {
			fmt.Printf("Go to Sleep for %s...\n\n", duration)
			time.Sleep(duration)
			fmt.Printf("\nLimiter length after Sleep : %d\n\n", len(burstyLimiter))
		}

		<-burstyLimiter
		fmt.Println("Decreased Limiter length   :", len(burstyLimiter))
		fmt.Printf("Request %2d                 : at %s\n\n", request, getMonotonicClock())
	}
}
