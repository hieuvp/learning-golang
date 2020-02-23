# Go Templates

## Table of Contents

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [Contextual Encoding](#contextual-encoding)
- [Template Actions](#template-actions)
- [Using Functions Inside Templates](#using-functions-inside-templates)
- [Creating the V in MVC](#creating-the-v-in-mvc)
- [References](#references)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Contextual Encoding

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=./hello.gohtml) -->
<!-- The below code snippet is automatically added from ./hello.gohtml -->

```gohtml
<h1>Hello, {{.Name}}!</h1>
```

<!-- AUTO-GENERATED-CONTENT:END -->

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=./hello.go) -->
<!-- The below code snippet is automatically added from ./hello.go -->

```go
package main

import (
	"html/template"
	"os"
)

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	data := struct {
		Name string
	}{"John Smith"}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
```

<!-- AUTO-GENERATED-CONTENT:END -->

```shell script
$ go run hello.go                                                                                                                                           1 ↵
<h1>Hello, John Smith!</h1>
```

## Template Actions

## Using Functions Inside Templates

## Creating the V in MVC

## References

- [An Introduction to Templates in Go](https://www.calhoun.io/intro-to-templates/)
- [ ] [Go Template Examples and Code Generator](https://github.com/phcollignon/Go-Template)
- [ ] [Using Go Templates](https://blog.gopheracademy.com/advent-2017/using-go-templates/)
