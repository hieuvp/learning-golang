# Go Templates

## Table of Contents

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [Contextual Encoding](#contextual-encoding)
  - [Creating a Template](#creating-a-template)
  - [A Contextual Encoding](#a-contextual-encoding)
- [Template Actions](#template-actions)
- [Using Functions Inside Templates](#using-functions-inside-templates)
- [Creating the V in MVC](#creating-the-v-in-mvc)
- [References](#references)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Contextual Encoding

### Creating a Template

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=./hello.gohtml) -->
<!-- The below code snippet is automatically added from ./hello.gohtml -->

```gohtml
<h1>Hello, {{.Name}}!</h1>
- https://github.com/yosssi/gohtml ==> formatter

- html formatter:
    - prettier
    - ...
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
	}{"Harrison Van"}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
```

<!-- AUTO-GENERATED-CONTENT:END -->

```shell script
go run hello.go > hello.html
```

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=./hello.html) -->
<!-- The below code snippet is automatically added from ./hello.html -->

```html
<h1>Hello, Harrison Van!</h1>
- https://github.com/yosssi/gohtml ==> formatter - html formatter: - prettier - ...
```

<!-- AUTO-GENERATED-CONTENT:END -->

### A Contextual Encoding

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=./context.gohtml) -->
<!-- The below code snippet is automatically added from ./context.gohtml -->

```gohtml
{{.Title}}
{{.HTML}}
{{.SafeHTML}}
{{.}}

<a title="{{.Title}}">
<a title="{{.HTML}}">

<a href="{{.HTML}}">
<a href="?q={{.HTML}}">
<a href="{{.Path}}">
<a href="?q={{.Path}}">

<!-- Encoding even works on non-string values! -->
<script>
    var dog = {{.Dog}};
    var map = {{.Map}};
    doWork({{.Title}});
</script>
```

<!-- AUTO-GENERATED-CONTENT:END -->

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=./context.go) -->
<!-- The below code snippet is automatically added from ./context.go -->

```go
package main

import (
	"html/template"
	"os"
)

// Test struct
type Test struct {
	HTML     string
	SafeHTML template.HTML
	Title    string
	Path     string
	Dog      Dog
	Map      map[string]string
}

// Dog struct
type Dog struct {
	Name string
	Age  int
}

func main() {
	t, err := template.ParseFiles("context.gohtml")
	if err != nil {
		panic(err)
	}

	data := Test{
		HTML:     "<h1>A header!</h1>",
		SafeHTML: template.HTML("<h1>A Safe header</h1>"),
		Title:    "Backslash! An in depth look at the \"\\\" character.",
		Path:     "/dashboard/settings",
		Dog:      Dog{"Fido", 6},
		Map: map[string]string{
			"key":       "value",
			"other_key": "other_value",
		},
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
```

<!-- AUTO-GENERATED-CONTENT:END -->

```shell script
go run context.go > context.html
```

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=./context.html) -->
<!-- The below code snippet is automatically added from ./context.html -->

```html
Backslash! An in depth look at the &#34;\&#34; character. &lt;h1&gt;A header!&lt;/h1&gt;
<h1>A Safe header</h1>
{&lt;h1&gt;A header!&lt;/h1&gt; &lt;h1&gt;A Safe header&lt;/h1&gt; Backslash! An in depth look
at the &#34;\&#34; character. /dashboard/settings {Fido 6} map[key:value other_key:other_value]}

<a title="Backslash! An in depth look at the &#34;\&#34; character.">
  <a title="&lt;h1&gt;A header!&lt;/h1&gt;">
    <a href="%3ch1%3eA%20header!%3c/h1%3e">
      <a href="?q=%3ch1%3eA%20header%21%3c%2fh1%3e">
        <a href="/dashboard/settings">
          <a href="?q=%2fdashboard%2fsettings">
            <script>
              var dog = { Name: "Fido", Age: 6 };
              var map = { key: "value", other_key: "other_value" };
              doWork('Backslash! An in depth look at the "\\" character.');
            </script></a
          ></a
        ></a
      ></a
    ></a
  ></a
>
```

<!-- AUTO-GENERATED-CONTENT:END -->

## Template Actions

## Using Functions Inside Templates

## Creating the V in MVC

## References

- [An Introduction to Templates in Go](https://www.calhoun.io/intro-to-templates/)
- [Go Template Examples and Code Generator](https://github.com/phcollignon/Go-Template)
- [ ] [Using Go Templates](https://blog.gopheracademy.com/advent-2017/using-go-templates/)
