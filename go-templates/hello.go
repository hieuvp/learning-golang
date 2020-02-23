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
		Title: "Users location",
		Users: []User{
			{
				Username: "Florin",
				Locations: map[string]Location{
					"Home": {
						Street:  "GoLand",
						ZipCode: "2018.3",
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
