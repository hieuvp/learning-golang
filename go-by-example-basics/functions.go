package main

import "fmt"

// A function that takes two integers and returns their sum as an integer
func plus(a int, b int) int {

	// Require explicit returns
	return a + b
}

// When you have multiple consecutive parameters of the same type
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {

	// Call a function just as you'd expect
	result := plus(1, 2)
	fmt.Println("1 + 2 =", result)

	result = plusPlus(1, 2, 3)
	fmt.Println("1 + 2 + 3 =", result)
}
