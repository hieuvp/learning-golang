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
$ go run hello-world.go
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
		i += 1
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

- There is no **Ternary Conditional Operator** (`condition ? exprIfTrue : exprIfFalse`) in Go, so you will need to use a full if statement even for basic conditions.


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
# Declare: [1 2 3 4 5]
# 2D Array: [[0 1 2] [1 2 3]]
```


## Slices

- Slices are a key data type in Go, giving a more powerful interface to sequences than arrays.

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=slices.go) -->
<!-- The below code snippet is automatically added from slices.go -->
```go
package main

import "fmt"

func main() {

	// Unlike arrays, slices are typed only by the elements they contain (not the number of elements)
	// To create an empty slice with non-zero length, use the built-in "make"
	// Here we make a slice of strings of length 3, initially zero-valued
	s := make([]string, 3)
	fmt.Println("Empty:", s)

	// Set and get just like with arrays
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("Set:", s)
	fmt.Println("Get:", s[2])

	// "len" returns the length of the slice
	fmt.Println("len:", len(s))

	// In addition to these basic operations, slices support several more that make them richer than arrays
	// One is the builtin append, which returns a slice containing one or more new values
	// Note that we need to accept a return value from append as we may get a new slice value
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// Slices can also be copy'd
	// Here we create an empty slice c of the same length as s and copy into c from s
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// Slices support a "slice" operator with the syntax slice[low:high]
	// For example, this gets a slice of the elements s[2], s[3], and s[4]
	l := s[2:5]
	fmt.Println("sl1:", l)

	// This slices up to (but excluding) s[5]
	l = s[:5]
	fmt.Println("sl2:", l)

	// And this slices up from (and including) s[2]
	l = s[2:]
	fmt.Println("sl3:", l)

	// We can declare and initialize a variable for slice in a single line as well
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

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
	fmt.Println("2d: ", twoD)
}
```
<!-- AUTO-GENERATED-CONTENT:END -->

```bash
$ go run slices.go

# emp: [  ]
# set: [a b c]
# get: c
# len: 3
# apd: [a b c d e f]
# cpy: [a b c d e f]
# sl1: [c d e]
# sl2: [a b c d e]
# sl3: [c d e f]
# dcl: [g h i]
# 2d:  [[0] [1 2] [2 3 4]]
```


## Maps

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=maps.go) -->
<!-- The below code snippet is automatically added from maps.go -->
```go
package main

import "fmt"

func main() {

	// To create an empty map, use the builtin make
	// make(map[key-type]val-type)
	m := make(map[string]int)

	// Set key/value pairs using typical name[key] = val syntax
	m["k1"] = 7
	m["k2"] = 13

	// Printing a map with e.g. fmt.Println will show all of its key/value pairs
	fmt.Println("map:", m)

	// Get a value for a key with name[key]
	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	// The builtin len returns the number of key/value pairs when called on a map
	fmt.Println("len:", len(m))

	// The builtin delete removes key/value pairs from a map
	delete(m, "k2")
	fmt.Println("map:", m)

	// The optional second return value when getting a value from a map indicates if the key was present in the map
	// This can be used to disambiguate between missing keys and keys with zero values like 0 or ""
	// Here we didn't need the value itself, so we ignored it with the blank identifier _
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// You can also declare and initialize a new map in the same line with this syntax
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
```
<!-- AUTO-GENERATED-CONTENT:END -->

```bash
$ go run maps.go

# map: map[k1:7 k2:13]
# v1:  7
# len: 2
# map: map[k1:7]
# prs: false
# map: map[bar:2 foo:1]
```


## Range

- `range` iterates over elements in a variety of data structures.

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=range.go) -->
<!-- The below code snippet is automatically added from range.go -->
```go
package main

import "fmt"

func main() {

	// Here we use range to sum the numbers in a slice
	// Arrays work like this too
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// range on arrays and slices provides both the index and value for each entry
	// Above we didn't need the index, so we ignored it with the blank identifier _
	// Sometimes we actually want the indexes though
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// range on map iterates over key/value pairs
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// range can also iterate over just the keys of a map
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// range on strings iterates over Unicode code points
	// The first value is the starting byte index of the rune and the second the rune itself
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
```
<!-- AUTO-GENERATED-CONTENT:END -->

```bash
$ go run range.go

# sum: 9
# index: 1
# a -> apple
# b -> banana
# key: a
# key: b
# 0 103
# 1 111
```


## Functions

## Multiple Return Values

## Variadic Functions

## Closures

## Recursion

## Pointers

## Structs

## Methods

## Interfaces

## Errors


## References

- [Go by Example](https://gobyexample.com/)
