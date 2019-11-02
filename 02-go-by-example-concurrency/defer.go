package main

import (
	"fmt"
	"os"
)

// Suppose we wanted to:
// 1. Create a file
// 2. Write to it
// 3. Then close when we are done
func main() {

	file := createFile("/tmp/defer.txt")

	// This will be executed at the end of the enclosing "func main()",
	// after "writeFile" has finished
	defer closeFile(file)

	writeFile(file)
}

func createFile(path string) *os.File {
	fmt.Println("Creating...")
	file, err := os.Create(path)

	if err != nil {
		panic(err)
	}
	return file
}

func writeFile(file *os.File) {
	fmt.Println("Writing...")
	fmt.Fprintln(file, "data")
}

func closeFile(file *os.File) {
	fmt.Println("Closing...")
	err := file.Close()

	// It is important to check for errors when closing a file,
	// even in a deferred function
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
