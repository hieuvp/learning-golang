package main

import (
	"fmt"
	"os"
)

// Suppose we wanted to create a file, write to it, and then close when we are done
// Here is how we could do that with "defer"
func main() {

	// Immediately after getting a file object with "createFile",
	// we "defer" the closing of that file with "closeFile"
	f := createFile("/tmp/defer.txt")

	// This will be executed at the end of the enclosing "func main()", after "writeFile" has finished
	defer closeFile(f)

	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("Creating...")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(file *os.File) {
	fmt.Println("Writing...")
	fmt.Fprintln(file, "data")
}

// It is important to check for errors when closing a file, even in a deferred function
func closeFile(file *os.File) {
	fmt.Println("Closing...")
	err := file.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
