package main

import (
	"html/template"
	"io/ioutil"
	"os"
)

// panic: template: UsersPage:10:21:
// executing "UsersPage" at <.username>: username
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
