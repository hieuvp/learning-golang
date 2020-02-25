# Go Templates

## Table of Contents

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [Hello World](#hello-world)
- [File Extensions](#file-extensions)
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

There two packages in go to work with templates:

- `text/templates` (Used for generating textual output)
- `html/templates` (Used for generating HTML output safe against code injection)

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

## File Extensions

```text
Supported by JetBrains plugin Go Template by default
├── example.gohtml
├── example.goyaml (not supported by default but a similar example)

Used by Helm Template
Used by PHP Smarty Template Engine (Supported by Smarty plugin in IntelliJ and PhpStorm)
Most prefer personally
├── example.tpl
├── example.html.tpl

Most used, even mentioned in
https://golang.org/pkg/html/template/
├── example.tmpl
├── example.html.tmpl
```

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

```text
Alternatively, type {{- /*gotype: */ -}},
place the caret after gotype:, press ⌃Space, and select the necessary type.
```

Now let’s start to add some output to our page so that we can deliver the data to it.
Normally you’d start typing something like `“<title>{{.”`
and expect the IDE to be smart enough and give you completion options
for the options after the dot.

This is where GoLand comes to help us.
We can now specify the type beforehand by
invoking the “Specify dot type” action via Alt + Enter and
select the type from the list of types available in the project.

<https://d3nmt5vlzunoa1.cloudfront.net/go/files/2018/12/Go-Template-optimized.gif>

This doesn’t work only for structure fields as the “Title” of the page works,
but it works for slices, slice elements, and even for elements that are
part of a map and are a more complex type.

<https://d3nmt5vlzunoa1.cloudfront.net/go/files/2018/12/Go-Template-2-optimized.gif>

Besides completion options,
once you specify the type of the dot in the template,
other functionality such as Navigate to Declaration, Find Usages,
or even Rename refactoring will work as the IDE has enough information to
complete these actions.

<https://d3nmt5vlzunoa1.cloudfront.net/go/files/2018/12/Go-Template-3-optimized.gif>

That's it for today.
We learned how we can get better code assistance
from the IDE when using the built-in Go template engine and
work with it more effectively.

Please let us know your feedback in the comments section below,
on Twitter, or on our issue tracker,
and tell us what would you like to learn more about in future articles.

## Contextual Encoding

## Template Actions

## Using Functions Inside Templates

## Creating the V in MVC

## References

- [Go Templates Made Easy](https://blog.jetbrains.com/go/2018/12/14/go-templates-made-easy/)
- [ ] [An Introduction to Templates in Go](https://www.calhoun.io/intro-to-templates/)
- [ ] [Go Template Examples and Code Generator](https://github.com/phcollignon/Go-Template)
