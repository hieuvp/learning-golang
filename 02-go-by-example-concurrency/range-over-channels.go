package main

import "fmt"

func main() {

	// We will iterate over "2" values in the "queue" channel
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// This range iterates over each element as it is received from "queue"
	// Because we closed the channel above,
	// the iteration terminates after receiving the "2" elements
	for element := range queue {
		fmt.Println(element)
	}
}
