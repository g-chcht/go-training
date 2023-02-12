package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%v %v\n", r.Method, r.URL)
	fmt.Fprintf(w, "Hello gophers!")
}

func search(w http.ResponseWriter, r *http.Request) {
	t := r.URL.Query().Get("t")
	p := r.URL.Query().Get("p")
	fmt.Printf("Param values t=%v, p=%v\n", t, p)

	fmt.Fprintf(w, "Searching for term=%v. Current page=%v", t, p)
}

func login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "login.html")
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() failed. err=%v", err)
			return
		}

		fmt.Fprintf(w, "Go login POST. value=%v\n", r.PostForm)
		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == "Go" && password == "rocks" {
			fmt.Fprintf(w, "You are now logged\n")
		} else {
			fmt.Fprintf(w, "Wrong username / password")
		}
	}
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/login", login)
	http.HandleFunc("/search", search)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
