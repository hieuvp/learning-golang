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

## If/Else

## Switch

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
