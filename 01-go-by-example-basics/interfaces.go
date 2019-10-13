package main

import (
	"fmt"
	"math"
)

// Basic "interface"
// Unlike "struct", instead of defining "fields" we define "methods"
type geometry interface {
	area() float64
	perimeter() float64
}

// We will implement "geometry" interface on "rect" and "circle" types
type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

// To implement an interface, we need to implement all the methods in the interface
// Here we implement "geometry" on "rect"
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perimeter() float64 {
	return 2*r.width + 2*r.height
}

// The implementation for "circle"
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// If a variable has an interface type,
// then we can call methods that are in the named interface
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// The "rect" and "circle" struct types both implement the "geometry" interface,
	// so we can use instances of these structs as arguments to "measure"
	measure(r)
	measure(c)
}
