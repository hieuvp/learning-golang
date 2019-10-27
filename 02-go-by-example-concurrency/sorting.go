package main

import (
	"fmt"
	"sort"
)

func main() {

	// Sort methods are specific to the builtin type
	// Note that sorting is in-place,
	// so it changes the given slice and does not return a new one

	// An example of sorting strings
	strings := []string{"c", "a", "b"}
	sort.Strings(strings)
	fmt.Println("strings  :", strings)

	// An example of sorting integers
	integers := []int{7, 2, 4}
	sort.Ints(integers)
	fmt.Println("integers :", integers)

	// We can also use "sort" to check if a slice is already in sorted order
	sorted := sort.IntsAreSorted(integers)
	fmt.Println("sorted   : ", sorted)
}
