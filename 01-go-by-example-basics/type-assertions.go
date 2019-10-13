package main

import "fmt"

func main() {
	var i interface{} = "hello"

	// Assert that the "interface value" "i" holds the "concrete type" "string",
	// and assign the underlying "string" value to the variable "s"
	s := i.(string)
	fmt.Println(s)

	// Test whether an interface value holds a specific type
	// An "ok" boolean value reports whether the assertion succeeded
	s, ok := i.(string)
	fmt.Println(s, ok)

	// If "ok" is "false"
	// then "f" will be zero-valued of type "float64" (in this case)
	// and no "panic" occurs
	f, ok := i.(float64)
	fmt.Println(f, ok)

	// "i" does not hold a "float64", the statement will trigger a "panic"
	f = i.(float64)
	fmt.Println(f)
}
