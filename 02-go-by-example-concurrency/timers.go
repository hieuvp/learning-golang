package main

import (
	"fmt"
	"time"
)

func main() {

	// "Timer" represents a single event in the future
	// You tell the "Timer" how long you want to wait,
	// and it provides a channel that will be notified at that time
	timer1 := time.NewTimer(2 * time.Second)

	// The "<-timer1.C" blocks on the timer's channel C
	// until it sends a value indicating that the timer expired
	<-timer1.C
	fmt.Println("Timer 1 expired")

	// If you just wanted to wait, you could have used "time.Sleep"
	// One reason a "Timer" may be useful is that
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
