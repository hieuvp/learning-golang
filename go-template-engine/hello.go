package main

import (
	"html/template"
	"io/ioutil"
	"os"
)

// panic: template: Page:10:21:
// executing "Page" at <.username>: username
// is an unexported field of struct type main.User

// It's not really intuitive,
// but templates (and encoding packages like JSON, for that matter)
// can't access unexported data members,
// so you have to export them somehow:

// “export” means “public” => with upper case first letter
// the reason is simple:
// the renderer package use the reflect package in order to get/set fields values
// the reflect package can only access public/exported struct fields.
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

	tpl, err := template.New("Page").Parse(string(message))
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
