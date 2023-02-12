package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	Country string `json:"country,omitempty"`
}

type User struct {
	Name     string  `json:"name"`
	Password string  `json:"-"`
	Email    string  `json:"email"`
	Ad       Address `json:address`
}

func users(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{Name: "Max",
			Password: "p1234",
			Email:    "max@gmail.com",
			Ad: Address{Street: "1 av mt royal",
				City:    "Montreal",
				Country: "Canada"}},
		{Name: "Leon",
			Password: "p123d",
			Email:    "leon@gmail.com",
			Ad: Address{Street: "50 blvd abcd",
				City: "Montreal"},
		},
	}
	b, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func main() {

	http.HandleFunc("/users", users)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
