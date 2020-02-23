# Go by Example - Basics

## Table of Contents

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [Hello World](#hello-world)
- [Values](#values)
- [Variables](#variables)
- [Constants](#constants)
- [For](#for)
- [If/Else](#ifelse)
- [Switch](#switch)
- [Arrays](#arrays)
- [Slices](#slices)
- [Maps](#maps)
- [Range](#range)
- [Functions](#functions)
- [Multiple Return Values](#multiple-return-values)
- [Variadic Functions](#variadic-functions)
- [Closures](#closures)
- [Recursion](#recursion)
- [Pointers](#pointers)
- [Structs](#structs)
- [Methods](#methods)
- [Interfaces](#interfaces)
- [Type Assertions](#type-assertions)
- [Errors](#errors)
- [References](#references)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Hello World

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=hello-world.go) -->
<!-- The below code snippet is automatically added from hello-world.go -->

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World")
}
```

<!-- AUTO-GENERATED-CONTENT:END -->

```bash
go run hello-world.go
```

```bash
# Build our program into a binary file
$ go build hello-world.go
$ ./hello-world
```

## Values

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=values.go) -->
<!-- The below code snippet is automatically added from values.go -->

```go
package main

import "fmt"

func main() {
	fmt.Println("go" + "lang")

	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
```

<!-- AUTO-GENERATED-CONTENT:END -->

```bash
$ go run values.go

# golang

# 1+1 = 2
# 7.0/3.0 = 2.3333333333333335

# false
# true
# false
```

## Variables

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=variables.go) -->
<!-- The below code snippet is automatically added from variables.go -->

```go
package main

import "fmt"

func main() {

	// "var" declares one or more variables
	var a = "initial"
	fmt.Println(a)

	// Declare multiple variables at once
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go will infer the type of initialized variables
	var d = true
	fmt.Println(d)

	// Variables declared without a corresponding initialization are zero-valued
	// e.g. the zero value for an "int" is "0"
	var e int
	fmt.Println(e)

	// The ":=" syntax is shorthand for declaring and initializing a variable
	// In this case: var f string = "apple"
	f := "apple"
	fmt.Println(f)
}
```

<!-- AUTO-GENERATED-CONTENT:END -->

```bash
$ go run variables.go

# initial
# 1 2
# true
# 0
# apple
```

## Constants

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=constants.go) -->
<!-- The below code snippet is automatically added from constants.go -->

```go
package main

import "fmt"

// "const" declares a constant value
const s string = "constant"

func main() {
	fmt.Println(s)

	// A "const" statement can appear anywhere a "var" statement can
	const n = 500000000

	// Constant expressions perform arithmetic with arbitrary precision
	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))
}
```

<!-- AUTO-GENERATED-CONTENT:END -->

```bash
$ go run constants.go

# constant
# 6e+11
# 600000000000
```

## For

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=for.go) -->
<!-- The below code snippet is automatically added from for.go -->

```go
package main

import "fmt"

func main() {

	// Single condition
	fmt.Println()
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	// A classic initial/condition/after for loop
	fmt.Println()
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// "for" without a condition will loop repeatedly
	// until you "break" out of the loop
	// or "return" from the enclosing function
	fmt.Println()
	for {
		fmt.Println("loop")
		break
	}

	// You can also "continue" to the next iteration of the loop
	fmt.Println()
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
```

<!-- AUTO-GENERATED-CONTENT:END -->

```bash
$ go run for.go

# 1
# 2
# 3

# 7
# 8
# 9

# loop

# 1
# 3
# 5
```

## If/Else

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=if-else.go) -->
<!-- The below code snippet is automatically added from if-else.go -->

```go
package main

import "fmt"

func main() {

	// You don't need parentheses around conditions in Go,
	// but that the braces are required
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// You can have an "if" statement without an "else"
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// A statement can precede conditionals;
	// Any variables declared in this statement are available in all branches
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
```

<!-- AUTO-GENERATED-CONTENT:END -->

```bash
$ go run if-else.go

# 7 is odd
# 8 is divisible by 4
# 9 has 1 digit
```

- There is no **Ternary Conditional Operator** (`condition ? exprIfTrue : exprIfFalse`) in Go,
  so you will need to use a full if statement even for basic conditions.

## Switch

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=switch.go) -->
<!-- The below code snippet is automatically added from switch.go -->

```go
package main

import (
	"fmt"
	"time"
)

func main() {

	// The basic "switch"
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// Use commas "," to separate multiple expressions in the same "case" statement
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	// Optional "default" case
	default:
		fmt.Println("It's a weekday")
	}

	// "switch" without an expression is an alternate way to express "if/else" logic
	t := time.Now()
	switch {
	// "case" expressions can be non-constants
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// A "type switch" compares types instead of values
	// Use this to discover the concrete type of an "interface value"
	whatAmI := func(i interface{}) {

		// Variable "t" will have the type corresponding to its clause
		// This syntax only works in "switch" statement
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
```

<!-- AUTO-GENERATED-CONTENT:END -->

```bash
$ go run switch.go

# Write 2 as two
# It's the weekend
# It's before noon

# I'm a bool
# I'm an int
# Don't know type string
```

## Arrays

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=arrays.go) -->
<!-- The below code snippet is automatically added from arrays.go -->

```go
package main

import "fmt"

func main() {

	// Create an array that will hold exactly 5 integers
	// By default, an array is zero-valued, which for integers means 0s
	var a [5]int
	fmt.Println("Empty:", a)

	// Set a value at an index
	a[4] = 100
	fmt.Println("Set:", a)

	// Get a value at an index
	fmt.Println("Get:", a[4])

	// The built-in "len" returns the length of an array
	fmt.Println("len:", len(a))

	// Declare and initialize an array in one line
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Init:", b)

	// Compose multi-dimensional data structures
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2D array:", twoD)
}
```

<!-- AUTO-GENERATED-CONTENT:END -->

```bash
$ go run arrays.go

# Empty: [0 0 0 0 0]
# Set: [0 0 0 0 100]
# Get: 100
# len: 5
# Init: [1 2 3 4 5]
# 2D array: [[0 1 2] [1 2 3]]
```

## Slices

> **Slices** are a key data type in Go, giving a more powerful interface to sequences than **Arrays**.

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=slices.go) -->
<!-- The below code snippet is automatically added from slices.go -->

```go
package main

import "fmt"

func main() {

	// To create an empty slice with non-zero length, use the built-in "make"
	// Here we make a slice of strings of length 3, initially zero-valued
	s := make([]string, 3)
	fmt.Println("Empty:", s)

	// Set and Get just like with arrays
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("Set:", s)
	fmt.Println("Get:", s[2])

	// "len" returns the length of the slice
	fmt.Println("len:", len(s))

	// The built-in "append", which returns a slice containing one or more new values
	// Note that we need to accept a return value from "append" as we may get a new slice value
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("Append:", s)
	fmt.Println("New length:", len(s))

	// Slices can also be copied
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("Copy:", c)

	// Support a "slice operator"
	// Syntax: slice[low:high]
	// e.g. To get a slice of the elements s[2], s[3], and s[4]
	fmt.Println("s[2:5] =", s[2:5])

	// Slice up to (but excluding) s[5]
	fmt.Println("s[:5] =", s[:5])

	// Slice up from (and including) s[2]
	fmt.Println("s[2:] =", s[2:])

	// Declare and initialize a variable for a slice in a single line
	// Just like an array, except we leave out the element count
	t := []string{"g", "h", "i"}
	fmt.Println("Init:", t)

	// Slices can be composed into multi-dimensional data structures
	// The length of the inner slices can vary, unlike with multi-dimensional arrays
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2D slice: ", twoD)
}
```

<!-- AUTO-GENERATED-CONTENT:END -->

```bash
$ go run slices.go

# Empty: [  ]
# Set: [a b c]
# Get: c
# len: 3

# Append: [a b c d e f]
# New length: 6

# Copy: [a b c d e f]

# s[2:5] = [c d e]
# s[:5] = [a b c d e]
# s[2:] = [c d e f]

# Init: [g h i]
# 2D slice:  [[0] [1 2] [2 3 4]]
```

## Maps

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=maps.go) -->
<!-- The below code snippet is automatically added from maps.go -->

```go
package main

import "fmt"

func main() {

	// To create an empty map, use the built-in "make"
	// Syntax: make(map[key-type]value-type)
	m := make(map[string]int)

	// Set key/value pairs
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("Map:", m)

	// Get a value for a key
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// The built-in "len" returns the number of key/value pairs
	fmt.Println("len:", len(m))

	// The built-in "delete" removes key/value pairs from a map
	delete(m, "k2")
	fmt.Println("Map:", m)

	// The optional second indicates if the key was present in the map
	// Use to disambiguate between missing keys and keys with zero values (e.g. 0 or "")
	// Here we didn't need the value itself, so we ignored it with the blank identifier "_"
	_, present := m["k2"]
	fmt.Println("Present:", present)

	// Declare and initialize a new map in the same line
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("Init:", n)
}
```

<!-- AUTO-GENERATED-CONTENT:END -->

```bash
$ go run maps.go

# Map: map[k1:7 k2:13]
# v1: 7
# len: 2
# Map: map[k1:7]
# Present: false
# Init: map[bar:2 foo:1]
```

## Range

> `range` iterates over elements in a variety of data structures.

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=range.go) -->
<!-- The below code snippet is automatically added from range.go -->

```go
package main

import "fmt"

func main() {

	// range on arrays and slices provides both the "index" and "value" for each entry
	nums := []int{2, 3, 4}
	for index, num := range nums {
		if num == 3 {
			fmt.Println("Index:", index)
		}
	}

	// Here we didn't need the "index", so we ignored it with the blank identifier "_"
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("Sum:", sum)

	// range on maps iterates over key/value pairs
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for key, value := range kvs {
		fmt.Printf("%s -> %s\n", key, value)
	}

	// range can also iterate over just the keys of a map
	for key := range kvs {
		fmt.Println("Key:", key)
	}

	// A Unicode code point is a unique number assigned to each Unicode character
	// range on strings iterates over Unicode code points
	// The first value is the starting byte index of the rune and the second the rune itself
	for index, character := range "go" {
		fmt.Println(index, character)
	}
}
```

<!-- AUTO-GENERATED-CONTENT:END -->

```bash
$ go run range.go

# Index: 1
# Sum: 9

# a -> apple
# b -> banana

# Key: a
# Key: b

# 0 103
# 1 111
```

## Functions

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=functions.go) -->
<!-- The below code snippet is automatically added from functions.go -->

```go
package main

import "fmt"

// A function that takes two integers and returns their sum as an integer
func plus(a int, b int) int {

	// Require explicit returns
	return a + b
}

// When you have multiple consecutive parameters of the same type
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {

	// Call a function just as you'd expect
	result := plus(1, 2)
	fmt.Println("1 + 2 =", result)

	result = plusPlus(1, 2, 3)
	fmt.Println("1 + 2 + 3 =", result)
}
```

<!-- AUTO-GENERATED-CONTENT:END -->

```bash
$ go run functions.go

# 1 + 2 = 3
# 1 + 2 + 3 = 6
```

## Multiple Return Values

> This feature is used often in idiomatic Go
> e.g. to return both **result** and **error** values from a function.

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=multiple-return-values.go) -->
<!-- The below code snippet is automatically added from multiple-return-values.go -->

```go
package main

import "fmt"

// "(int, int)" in this function signature shows that
// the function returns 2 integers
func values() (int, int) {
	return 3, 7
}

func main() {

	// 2 different return values from the call with multiple assignment
	a, b := values()
	fmt.Println(a)
	fmt.Println(b)

	// Use the blank identifier "_",
	// if you only want a subset of the returned values
	_, c := values()
	fmt.Println(c)
}
```

<!-- AUTO-GENERATED-CONTENT:END -->

```bash
$ go run multiple-return-values.go

# 3
# 7

# 7
```

## Variadic Functions

> Variadic functions can be called with any number of trailing arguments (e.g. `fmt.Println`).

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=variadic-functions.go) -->
<!-- The below code snippet is automatically added from variadic-functions.go -->

```go
package main

import "fmt"

// Here's a function that will take an arbitrary number of integers as arguments
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	// Variadic functions can be called in the usual way with individual arguments
	sum(1, 2)
	sum(1, 2, 3)

	// If you already have multiple arguments in a slice
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
```

<!-- AUTO-GENERATED-CONTENT:END -->

```bash
$ go run variadic-functions.go

# [1 2] 3
# [1 2 3] 6

# [1 2 3 4] 10
```

## Closures

> Go supports **anonymous functions**, which can form **closures**.
> **Anonymous functions** are useful
> when you want to define a function inline without having to name it.

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=closures.go) -->
<!-- The below code snippet is automatically added from closures.go -->

```go
package main

import "fmt"

// This function intSeq returns another function
func intSeq() func() int {
	i := 0

	// We define anonymously
	return func() int {
		i++
		return i
	}
}

func main() {

	// We call intSeq, assigning the result (a function) to nextInt
	nextInt := intSeq()

	// See the effect of the closure by calling nextInt a few times
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// Create and test a new one
	newInt := intSeq()
	fmt.Println(newInt())
}
```

<!-- AUTO-GENERATED-CONTENT:END -->

```bash
$ go run closures.go

# 1
# 2
# 3

# 1
```

## Recursion

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=recursion.go) -->
<!-- The below code snippet is automatically added from recursion.go -->

```go
package main

import "fmt"

// This "fact" function calls itself until it reaches the base case of "fact(0)"
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))
}
```

<!-- AUTO-GENERATED-CONTENT:END -->

```bash
$ go run recursion.go

# 5040
```

## Pointers

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=pointers.go) -->
<!-- The below code snippet is automatically added from pointers.go -->

```go
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
```

<!-- AUTO-GENERATED-CONTENT:END -->

```bash
$ go run pointers.go

# Initial: 1
# zeroValue: 1

# Pointer: 0xc00009a000
# iPointer = 0xc00009a000
# zeroPointer: 0
```

## Structs

> **Structs** are typed collections of **fields**,
> they're useful for grouping data together to form records.

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=structs.go) -->
<!-- The below code snippet is automatically added from structs.go -->

```go
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
```

<!-- AUTO-GENERATED-CONTENT:END -->

```bash
$ go run structs.go

# {Bob 20}
# {Alice 30}
# {Fred 0}
# &{Ann 40}
# &{Jon 42}
# Sean

# 50
# 50

# &{Sean 51}
```

## Methods

> Go supports **methods** defined on **struct types**.

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=methods.go) -->
<!-- The below code snippet is automatically added from methods.go -->

```go
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
```

<!-- AUTO-GENERATED-CONTENT:END -->

```bash
$ go run methods.go

# area: 50
# perimeter: 30

# area: 50
# perimeter: 30
```

## Interfaces

> **Interfaces** are named collections of **method** signatures.

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=interfaces.go) -->
<!-- The below code snippet is automatically added from interfaces.go -->

```go
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
```

<!-- AUTO-GENERATED-CONTENT:END -->

```bash
$ go run interfaces.go

# {3 4}
# 12
# 14

# {5}
# 78.53981633974483
# 31.41592653589793
```

## Type Assertions

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=type-assertions.go) -->
<!-- The below code snippet is automatically added from type-assertions.go -->

```go
package main

import "fmt"

func main() {
	var i interface{} = "hello"

	// Assert that the "interface value" "i" holds the "concrete type" "string"
	// Assign the underlying "string" value to the variable "s"
	s := i.(string)
	fmt.Println(s)

	// To test whether an interface value holds a specific type
	// An "ok" boolean value reports whether the assertion succeeded
	s, ok := i.(string)
	fmt.Println(s, ok)

	// Because "ok" results in "false",
	// "f" will be zero-valued of type "float64" and no "panic" occurs
	f, ok := i.(float64)
	fmt.Println(f, ok)

	// Because "i" does not hold a "float64", so this statement will trigger a "panic"
	f = i.(float64)
	fmt.Println(f)
}
```

<!-- AUTO-GENERATED-CONTENT:END -->

```bash
$ go run type-assertions.go

# hello
# hello true

# 0 false
# panic: interface conversion: interface {} is string, not float64
```

## Errors

> In Go, it is idiomatic to communicate **errors** via an explicit, separate **return value**.

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=errors.go) -->
<!-- The below code snippet is automatically added from errors.go -->

```go
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
```

<!-- AUTO-GENERATED-CONTENT:END -->

```bash
$ go run errors.go

# f1 worked: 10
# f1 failed: cannot work with 42

# f2 worked: 10
# f2 failed: 42 - cannot work with it

# 42
# cannot work with it
```

- This approach makes it easy to see which functions return errors
  and to handle them using the same language
  constructs employed for any other, non-error tasks.

## References

- [Go by Example](https://gobyexample.com/)
