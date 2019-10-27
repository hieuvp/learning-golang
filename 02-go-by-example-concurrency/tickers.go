package main

import (
	"fmt"
	"time"
)

func main() {

	// "Ticker" uses a similar mechanism to "Timer": a channel that is sent values
	// Here we await the values as they arrive "every 500ms"
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at :", t)
			}
		}
	}()

	// "Ticker" can be stopped like "Timer"
	// Once a "Ticker" is stopped, it won't receive any more values on its channel
	time.Sleep(1600 * time.Millisecond)

	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")

	time.Sleep(2000 * time.Millisecond)
}
