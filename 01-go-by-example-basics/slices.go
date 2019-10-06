package main

import "fmt"

func main() {

	// To create an empty slice with non-zero length, use the built-in "make"
	// Here we make a slice of strings of length 3, initially zero-valued
	s := make([]string, 3)
	fmt.Println("Empty:", s)

	// Set and Get just like with arrays
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("Set:", s)
	fmt.Println("Get:", s[2])

	// "len" returns the length of the slice
	fmt.Println("len:", len(s))

	// The built-in "append", which returns a slice containing one or more new values
	// Note that we need to accept a return value from "append" as we may get a new slice value
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("Append:", s)
	fmt.Println("New length:", len(s))

	// Slices can also be copied
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("Copy:", c)

	// Support a "slice operator"
	// Syntax: slice[low:high]
	// e.g. To get a slice of the elements s[2], s[3], and s[4]
	fmt.Println("s[2:5] =", s[2:5])

	// Slice up to (but excluding) s[5]
	fmt.Println("s[:5] =", s[:5])

	// Slice up from (and including) s[2]
	fmt.Println("s[2:] =", s[2:])

	// Declare and initialize a variable for a slice in a single line
	// Just like an array, except we leave out the element count
	t := []string{"g", "h", "i"}
	fmt.Println("Init:", t)

	// Slices can be composed into multi-dimensional data structures
	// The length of the inner slices can vary, unlike with multi-dimensional arrays
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2D slice: ", twoD)
}
