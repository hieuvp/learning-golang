package main

import "fmt"

func main() {

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"

	// Here we "close" the channel,
	// the iteration will terminate after receiving "2" elements
	close(queue)

	// This "range" iterates over each element as it is received from "queue"
	for element := range queue {
		fmt.Println(element)
	}
}
