package main

import "fmt"

func main() {
	var i interface{} = "hello"

	// Assert that the "interface value" "i" holds the "concrete type" "string"
	// Assign the underlying "string" value to the variable "s"
	s := i.(string)
	fmt.Println(s)

	// To test whether an interface value holds a specific type
	// An "ok" boolean value reports whether the assertion succeeded
	s, ok := i.(string)
	fmt.Println(s, ok)

	// Because "ok" results in "false",
	// "f" will be zero-valued of type "float64" and no "panic" occurs
	f, ok := i.(float64)
	fmt.Println(f, ok)

	// Because "i" does not hold a "float64", so this statement will trigger a "panic"
	f = i.(float64)
	fmt.Println(f)
}
