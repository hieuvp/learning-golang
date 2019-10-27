package main

import (
	"fmt"
	"time"
)

func main() {

	// "Timer" represents a single event in the future
	// You tell the "Timer" how long you want to wait
	timer1 := time.NewTimer(2 * time.Second)

	// It provides a channel that will send a value indicating when the "Timer" expired
	<-timer1.C
	fmt.Println("Timer 1 expired")

	// If you just wanted to wait, you could have used "time.Sleep"
	// One reason a "Timer" maybe useful is that,
	// you can "stop" the "Timer" before it expires
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}
