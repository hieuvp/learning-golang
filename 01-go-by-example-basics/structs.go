package main

import "fmt"

// Basic "struct"
type person struct {
	name string
	age  int
}

// "NewPerson" constructs a new "person" struct with the given "name"
func NewPerson(name string) *person {
	p := person{name: name}
	p.age = 42

	// Return a pointer to local variable safely
	return &p
}

func main() {

	// Create a new struct
	fmt.Println(person{"Bob", 20})

	// Name the fields when initializing a struct
	fmt.Println(person{age: 30, name: "Alice"})

	// Omitted fields will be zero-valued
	fmt.Println(person{name: "Fred"})

	// An "&" prefix yields a pointer to the struct
	fmt.Println(&person{name: "Ann", age: 40})

	// It is idiomatic to encapsulate new struct creation in "constructor" functions
	fmt.Println(NewPerson("Jon"))

	// Access struct fields with a dot (".")
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// When using "." with a struct pointer,
	// the pointer will be "dereferenced automatically"
	sp := &s
	fmt.Println(sp.age)
	fmt.Println((*sp).age)

	// Structs are mutable
	sp.age = 51
	fmt.Println(sp)
}
