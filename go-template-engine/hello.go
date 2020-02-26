package main

import (
	"html/template"
	"io/ioutil"
	"os"
)

// Templates (or Encoding Packages like JSON, YAML, Viper,...)
// cannot access "unexported" data members, so we have to export them somehow.

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

	// Page struct
	Page struct {
		Title string
		Users []User
	}
)

func main() {
	message, err := ioutil.ReadFile("hello.gohtml")
	if err != nil {
		panic(err)
	}

	// Parsing Templates
	//
	// The three most important and most frequently used functions are:
	//
	// New — allocates new, undefined template,
	// Parse — parses given template string and return parsed template,
	// Execute — applies parsed template to the data structure and writes result to the given writer.

	// Go standard library function template.New allocate a new undefined template with associated name.

	// The name of the template–unsurprisingly–is to name the template.

	// What is it good for?
	// As long as you don't want to refer to the template, it doesn't really matter.
	// But if you want to refer to it, then yes, you refer to it by its name.

	// When would you want to refer to it?\
	// When you want to include a template in another
	// e.g. using the {{template}} action,
	// or when you want to execute a specific template
	// using Template.ExecuteTemplate().

	// error checks omitted for brevity
	// Executes default, "one":
	// t.Execute(os.Stdout, nil)

	// Executes explicit, "one":
	// t.ExecuteTemplate(os.Stdout, "one", nil)

	// Executes explicit, "other":
	// t.ExecuteTemplate(os.Stdout, "other", nil)

	tpl, err := template.New("hello").Parse(string(message))
	if err != nil {
		panic(err)
	}

	page := Page{
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
			{
				Username: "Hieu Van",
				Locations: map[string]Location{
					"Home": {
						Street:  "GitHub",
						ZipCode: "2020",
					},
				},
			},
		},
	}

	err = tpl.Execute(os.Stdout, page)
	if err != nil {
		panic(err)
	}
}
