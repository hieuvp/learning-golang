package main

import "fmt"

// "zeroval" has an "int" parameter, so arguments will be passed to it by value
// "zeroval" will get a copy of "ival" distinct from the one in the calling function
func zeroval(ival int) {
	ival = 0
}

// "zeroptr" in contrast has an "*int" parameter, meaning that it takes an int pointer
func zeroptr(iptr *int) {
	// The "*iptr" code dereferences the pointer from its memory address to the current value at that address
	// Assigning a value to a dereferenced pointer changes the value at the referenced address
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("Initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// The "&i" syntax gives the memory address of i (a.k.a. "a pointer to i")
	fmt.Println("Pointer:", &i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)
}
