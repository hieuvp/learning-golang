package main

import "fmt"

func main() {

	// To create an empty map, use the built-in "make"
	// Syntax: make(map[key-type]value-type)
	m := make(map[string]int)

	// Set key/value pairs
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("Map:", m)

	// Get a value for a key
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// The built-in "len" returns the number of key/value pairs
	fmt.Println("len:", len(m))

	// The built-in "delete" removes key/value pairs from a map
	delete(m, "k2")
	fmt.Println("Map:", m)

	// The optional second indicates if the key was present in the map
	// Use to disambiguate between missing keys and keys with zero values (e.g. 0 or "")
	// Here we didn't need the value itself, so we ignored it with the blank identifier "_"
	_, present := m["k2"]
	fmt.Println("Present:", present)

	// Declare and initialize a new map in the same line
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("Init:", n)
}
