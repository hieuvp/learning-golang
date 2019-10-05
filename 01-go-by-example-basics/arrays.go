package main

import "fmt"

func main() {

	// Create an array that will hold exactly 5 integers
	// By default, an array is zero-valued, which for integers means 0s
	var a [5]int
	fmt.Println("Empty:", a)

	// Set a value at an index
	a[4] = 100
	fmt.Println("Set:", a)

	// Get a value at an index
	fmt.Println("Get:", a[4])

	// The built-in "len" returns the length of an array
	fmt.Println("len:", len(a))

	// Declare and initialize an array in one line
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Init:", b)

	// Compose multi-dimensional data structures
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2D array:", twoD)
}
