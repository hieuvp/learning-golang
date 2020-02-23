package main

import (
	"fmt"
	"reflect"
)

func main() {

	// Here we make a channel of strings buffering up to "2" values
	messages := make(chan string, 2)
	fmt.Println(reflect.TypeOf(messages))

	// Because this channel is buffered,
	// we can send these values into the channel without a corresponding concurrent receive
	messages <- "buffered"
	messages <- "channel"

	// Later we can receive these "2" values as usual
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
