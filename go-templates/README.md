# Go Templates

## Table of Contents

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [Hello World](#hello-world)
- [Contextual Encoding](#contextual-encoding)
- [Template Actions](#template-actions)
- [Using Functions Inside Templates](#using-functions-inside-templates)
- [Creating the V in MVC](#creating-the-v-in-mvc)
- [References](#references)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## Hello World

Go ships with its own template engine, split into two packages,
`text/template` and `html/template`.
These packages are similar in functionality,
with the difference that `html/template` allows
a user to generate HTML code that is safe against code injection,
making it suitable for use on web pages and emails.

To use code assistance in Go templates,
ensure that your project folder is under GOPATH (for example, go/src/myProject).
To check your GOPATH, open settings ⌘, and navigate to Go | GOPATH.

<https://www.jetbrains.com/help/go/configuring-goroot-and-gopath.html#gopath>
<https://www.jetbrains.com/help/go/integration-with-go-templates.html>
<https://github.com/apronichev/documentation-code-examples/tree/master/goTemplates>

Go supports a much better way to manage your project using Go Modules.
Follow this article to know more. Go workspaces (GOPATH) will be deprecated soon.

## Contextual Encoding

## Template Actions

## Using Functions Inside Templates

## Creating the V in MVC

## References

- [ ] [Go Templates Made Easy](https://blog.jetbrains.com/go/2018/12/14/go-templates-made-easy/)
- [ ] [An Introduction to Templates in Go](https://www.calhoun.io/intro-to-templates/)
- [ ] [Go Template Examples and Code Generator](https://github.com/phcollignon/Go-Template)
