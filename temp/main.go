package main

import (
	"html/template"
	"net/http"
)

type Person struct {
	Name             string
	Age              int
	Department       string
	Favoritelanguage string
}

func handler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	person := Person{
		Name:             "Stephen",
		Age:              33,
		Department:       "AI/ML",
		Favoritelanguage: "idoma",
	}
	err = tmpl.Execute(w, person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", handler)

	http.ListenAndServe(":8088", nil)
}
