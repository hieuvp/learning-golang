package main

import "fmt"

func main() {

	// To create an empty slice with non-zero length, use the built-in "make"
	// Here we make a slice of strings of length 3, initially zero-valued
	s := make([]string, 3)
	fmt.Println("Empty:", s)

	// Set and get just like with arrays
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("Set:", s)
	fmt.Println("Get:", s[2])

	// "len" returns the length of the slice
	fmt.Println("len:", len(s))

	// Unlike arrays, slices can be resized using the built-in append function
	// In addition to these basic operations, slices support several more that make them richer than arrays
	// One is the builtin append, which returns a slice containing one or more new values
	// Note that we need to accept a return value from append as we may get a new slice value
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("Append:", s)
	fmt.Println("len:", len(s))

	// Slices can also be copy'd
	// Here we create an empty slice c of the same length as s and copy into c from s
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("Copy:", c)

	// Slices support a "slice" operator with the syntax slice[low:high]
	// For example, this gets a slice of the elements s[2], s[3], and s[4]
	l := s[2:5]
	fmt.Println("sl1:", l)

	// This slices up to (but excluding) s[5]
	l = s[:5]
	fmt.Println("sl2:", l)

	// And this slices up from (and including) s[2]
	l = s[2:]
	fmt.Println("sl3:", l)

	// Declare and initialize a variable for slice in a single line
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

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
