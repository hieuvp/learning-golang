package main

import (
	"errors"
	"fmt"
)

// By convention, errors are the "last return value"
// and have type "error", a built-in interface
func f1(arg int) (int, error) {
	if arg == 42 {

		// "errors.New" constructs a basic error value with the given error message
		return -1, errors.New("cannot work with 42")
	}

	// A "nil" value in the error position indicates that there was no error
	return arg + 3, nil
}

// It is possible to use custom types as errors
// by implementing an "Error()" method to conform the "error" interface
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "cannot work with it"}
	}

	return arg + 3, nil
}

func main() {

	for _, i := range []int{7, 42} {
		// The use of an inline error check is a common idiom in Go code
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	// Get the error as an instance of the custom error type via "type assertion"
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
