package main

import "fmt"

// Person struct
type Person struct {
	name string
	age  int
}

// NewPerson constructs a new Person struct with the given "name"
func NewPerson(name string) *Person {
	p := Person{name: name}
	p.age = 42

	// Return a pointer to local variable safely
	return &p
}

func main() {

	// Create a new struct
	fmt.Println(Person{"Bob", 20})

	// Name the fields when initializing a struct
	fmt.Println(Person{age: 30, name: "Alice"})

	// Omitted fields will be zero-valued
	fmt.Println(Person{name: "Fred"})

	// An "&" prefix yields a pointer to the struct
	fmt.Println(&Person{name: "Ann", age: 40})

	// It is idiomatic to encapsulate new struct creation in "constructor" functions
	fmt.Println(NewPerson("Jon"))

	// Access struct fields with a dot (".")
	s := Person{name: "Sean", age: 50}
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
