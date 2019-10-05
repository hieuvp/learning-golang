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

## Slices

## Maps

## Range

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
