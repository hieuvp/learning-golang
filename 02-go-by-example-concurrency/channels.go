package main

import (
	"fmt"
	"reflect"
)

func main() {

	// Create a new channel with "make(chan value-type)"
	// Channels are typed by the values they convey
	messages := make(chan string)
	fmt.Println(reflect.TypeOf(messages))

	// "channel <-": send a value into a channel
	go func() { messages <- "ping" }()

	// "<-channel": receive a value from the channel
	msg := <-messages
	fmt.Println(msg)
}
