package main

import "fmt"

// "zeroValue" has an "int" parameter, so arguments will be passed into it by "int value"
// "zeroValue" will get a copy of "iValue" distinct from the one in the calling function
func zeroValue(iValue int) {
	iValue = 0
}

// "zeroPointer" in contrast has an "*int" parameter, meaning that it takes an "int pointer"
func zeroPointer(iPointer *int) {
	fmt.Println("iPointer =", iPointer)

	// "*iPointer": dereference the pointer from its "memory address"
	// to the current value at that address
	// Assign a "value" to a "dereferenced pointer" changes the "value" at the "referenced address"
	*iPointer = 0
}

func main() {
	i := 1
	fmt.Println("Initial:", i)

	zeroValue(i)
	fmt.Println("zeroValue:", i)

	// The "&i" syntax gives the "memory address" of "i" (a.k.a. "a pointer to i")
	fmt.Println("Pointer:", &i)

	zeroPointer(&i)
	fmt.Println("zeroPointer:", i)
}
