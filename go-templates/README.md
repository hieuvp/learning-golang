# Go Templates

## Table of Contents

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [Hello World](#hello-world)
- [JetBrains Code Assistance](#jetbrains-code-assistance)
- [Contextual Encoding](#contextual-encoding)
- [Template Actions](#template-actions)
- [Using Functions Inside Templates](#using-functions-inside-templates)
- [Creating the V in MVC](#creating-the-v-in-mvc)
- [References](#references)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Hello World

- Go ships with its own template engine, split into two packages,
  `text/template` and `html/template`.
- These packages are similar in functionality,
  with the difference that `html/template` allows
  a user to generate HTML code that is safe against code injection,
  making it suitable for use on web pages and emails.

<br />

**`hello.go`**

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=./hello.go) -->
<!-- The below code snippet is automatically added from ./hello.go -->

```go
package main

import (
	"io/ioutil"
	"os"
	"text/template"
)

type (
	// Location struct
	Location struct {
		Street  string
		ZipCode string
	}

	// User struct
	User struct {
		Username  string
		Locations map[string]Location
	}

	// UsersPage struct
	UsersPage struct {
		Title string
		Users []User
	}
)

func main() {
	message, err := ioutil.ReadFile("hello.gohtml")
	if err != nil {
		panic(err)
	}

	t, err := template.New("UsersPage").Parse(string(message))
	if err != nil {
		panic(err)
	}

	p := UsersPage{
		Title: "Users Location",
		Users: []User{
			{
				Username: "Harrison Van",
				Locations: map[string]Location{
					"Home": {
						Street:  "ShopBack",
						ZipCode: "2020",
					},
				},
			},
		},
	}

	err = t.Execute(os.Stdout, p)
	if err != nil {
		panic(err)
	}
}
```

<!-- AUTO-GENERATED-CONTENT:END -->

<br />

**`hello.gohtml`**

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=./hello.gohtml) -->
<!-- The below code snippet is automatically added from ./hello.gohtml -->

```gohtml
{{- /*gotype: templating.UsersPage*/ -}}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>{{.Title}}</title>
</head>
<body>
{{range .Users}}
    <div>Username: {{.Username}} </div>
    <div>Location: {{.Locations.Home.Street}}</div>
{{else}}
    <div>No users found</div>
{{end}}
</body>
</html>
```

<!-- AUTO-GENERATED-CONTENT:END -->

<br />

```shell script
go run hello.go
```

<!-- AUTO-GENERATED-CONTENT:START (CODE:src=./output/hello.html) -->
<!-- The below code snippet is automatically added from ./output/hello.html -->

```html
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <title>Users Location</title>
  </head>
  <body>
    <div>Username: Harrison Van</div>
    <div>Location: ShopBack</div>
  </body>
</html>
```

<!-- AUTO-GENERATED-CONTENT:END -->

## JetBrains Code Assistance

To use code assistance in Go templates:

- Ensure that your project folder is under `GOPATH`.
  To check your `GOPATH`, open settings `⌘`, and navigate to `Go | GOPATH`.

- However, `GOPATH` will be deprecated soon.
  As of `1.11`, Go supports a much better way to manage your project using `Go Modules`.
  It enables the use of modules when the current directory
  or any parent directory has a `go.mod`, provided the directory is outside `GOPATH`.

  ```shell script
  $ go mod init templating
  go: creating new go.mod: module templating
  ```

- <https://www.jetbrains.com/help/go/configuring-goroot-and-gopath.html#gopath>
- <https://www.jetbrains.com/help/go/integration-with-go-templates.html>

Place the caret at the head tag and press Alt+Enter, select Add 'title'
Inside the title tag, type {\{.}}
Place the caret inside {\{.}} and press Alt+Enter, select Specify dot type
In auto-completion popup, select goTemplates.UsersPage
Inside {\{.}}, press Ctrl+Space and select Title
You can use tpl_example.gohtml as a reference

## Contextual Encoding

## Template Actions

## Using Functions Inside Templates

## Creating the V in MVC

## References

- [Go Templates Made Easy](https://blog.jetbrains.com/go/2018/12/14/go-templates-made-easy/)
- [ ] [An Introduction to Templates in Go](https://www.calhoun.io/intro-to-templates/)
- [ ] [Go Template Examples and Code Generator](https://github.com/phcollignon/Go-Template)
