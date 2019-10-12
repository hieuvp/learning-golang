package main

import "fmt"

type rectangle struct {
	width, height int
}

// Compare to a "function", a "method" is declared by additionally specifying the "receiver"
// This "area" method has a "receiver type" of "*rectangle"
func (r *rectangle) area() int {
	return r.width * r.height
}

// Methods can be defined for either "pointer" or "value" receiver types
func (r rectangle) perimeter() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rectangle{width: 10, height: 5}

	// Here we call the 2 methods defined for our struct
	fmt.Println("area:", r.area())
	fmt.Println("perimeter:", r.perimeter())

	// For method calls, Go automatically handles conversion between "values" and "pointers"
	// Use a "pointer" "receiver type"
	// to avoid copying on method calls or
	// to allow the method to mutate the receiving struct
	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perimeter:", rp.perimeter())
}
