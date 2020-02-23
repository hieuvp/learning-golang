package main

import (
	"fmt"
	"sort"
)

// Sometimes, we will want to sort a collection by something other than its natural order
// e.g. to sort strings by their length instead of alphabetically

// Here is an example of custom "sort.Sort", "func Sort(data Interface)"

// Create a "byLength" type that is an "alias" for the built-in "[]string" type
type byLength []string

// We implement "sort.Interface"
// - Len() int: the number of elements in the collection
// - Less(i, j int) bool: report whether the element with index "i" should sort before the element with index "j"
// - Swap(i, j int): swap the elements with indexes "i" and "j"
func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Less(i, j int) bool {
	// In our case,
	// we want to sort in order of increasing string length
	return len(s[i]) < len(s[j])
}
func (s byLength) Swap(i, j int) {
	// Reassign values
	s[i], s[j] = s[j], s[i]
}

func main() {

	fruits := []string{"peach", "banana", "kiwi"}

	// Convert the original "fruits" "slice" to "byLength",
	// and then use the "Sort" function
	sort.Sort(byLength(fruits))

	fmt.Println(fruits)
}
