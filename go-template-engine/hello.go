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
