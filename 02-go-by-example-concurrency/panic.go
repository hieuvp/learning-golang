package main

import "os"

func main() {
	// Syntax
	panic("a problem")

	// A common use of "panic" is
	// to abort if a function returns an error value
	// that we do not know how to handle or do not want to handle

	// Panicking if we get an unexpected error when creating a new file
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
