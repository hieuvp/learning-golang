package main

import (
	"fmt"
	"sort"
)

func main() {

	// Sort methods are specific to the built-in type
	// Note that, the sorting changes the given slice and does not return a new one

	// Sorting strings
	strings := []string{"c", "a", "b"}
	sort.Strings(strings)
	fmt.Println("strings  :", strings)

	// Sorting integers
	integers := []int{7, 2, 4}
	sort.Ints(integers)
	fmt.Println("integers :", integers)

	// Check if a slice is already in sorted order
	sorted := sort.IntsAreSorted(integers)
	fmt.Println("sorted   : ", sorted)
}
