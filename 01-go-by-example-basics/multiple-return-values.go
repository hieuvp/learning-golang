package main

import "fmt"

// "(int, int)" in this function signature shows that
// the function returns 2 integers
func values() (int, int) {
	return 3, 7
}

func main() {

	// 2 different return values from the call with multiple assignment
	a, b := values()
	fmt.Println(a)
	fmt.Println(b)

	// Use the blank identifier "_",
	// if you only want a subset of the returned values
	_, c := values()
	fmt.Println(c)
}
