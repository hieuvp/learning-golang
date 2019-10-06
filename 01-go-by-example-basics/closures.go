package main

import "fmt"

// This function intSeq returns another function
func intSeq() func() int {
	i := 0

	// We define anonymously
	return func() int {
		i++
		return i
	}
}

func main() {

	// We call intSeq, assigning the result (a function) to nextInt
	nextInt := intSeq()

	// See the effect of the closure by calling nextInt a few times
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// Create and test a new one
	newInt := intSeq()
	fmt.Println(newInt())
}
