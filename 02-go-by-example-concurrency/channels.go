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

	// Send a value into a channel using the channel "<-" syntax
	// Here we send "ping" to the messages channel we made above, from a new goroutine
	go func() { messages <- "ping" }()

	// The "<-channel" syntax receives a value from the channel
	// Here we will receive the "ping" message we sent above and print it out
	msg := <-messages
	fmt.Println(msg)
}
