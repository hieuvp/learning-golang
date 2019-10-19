package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// Here is a "non-blocking receive"
	// If a value is available on "messages" then "select" will take the "<-messages" case
	// If not it will "immediately" take the "default" case
	select {
	case message := <-messages:
		fmt.Println("received message :", message)
	default:
		fmt.Println("no message received")
	}

	// A "non-blocking send" works similarly
	// "message" cannot be sent to the "messages" channel,
	// because the channel has no buffer and there is no receiver
	// Therefore the "default" case is selected
	message := "hi"
	select {
	case messages <- message:
		fmt.Println("sent message :", message)
	default:
		fmt.Println("no message sent")
	}

	// Here we attempt "non-blocking receives" on both "messages" and "signals"
	select {
	case message := <-messages:
		fmt.Println("received message :", message)
	case signal := <-signals:
		fmt.Println("received signal :", signal)
	default:
		fmt.Println("no activity")
	}
}
