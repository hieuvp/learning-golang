package main

import (
	"fmt"
	"time"
)

func main() {

	// We will "select" across two channels
	c1 := make(chan string)
	c2 := make(chan string)

	// Each channel will receive a value after some amount of time
	go func() {
		time.Sleep(4 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// Use "select" to await both of these values simultaneously
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Received :", msg1)
		case msg2 := <-c2:
			fmt.Println("Received :", msg2)
		}
	}
}
