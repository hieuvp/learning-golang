package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 4; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// Running synchronously
	f("Direct")

	// To invoke this function in a goroutine, use "go f(s)"
	// This new goroutine will execute concurrently with the calling one
	go f("Goroutine")

	// You can also start a goroutine for an anonymous function call
	go func(msg string) {
		fmt.Println(msg)
	}("Going")

	// Our two function calls above are running "asynchronously" in separate goroutines now

	// Pause the current goroutine for at least the "Duration"
	fmt.Println("Sleep with Duration:", time.Second)
	time.Sleep(time.Second)
	fmt.Println("Done")
}
