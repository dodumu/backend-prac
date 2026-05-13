package main

import (
	"html/template"
	"net/http"
	"strconv"
)

type Student struct {
	Name             string
	Age              int
	Email            string
	Department       string
	FavoriteLanguage string
}

func handler(w http.ResponseWriter, r *http.Request) {

	// GET REQUEST
	if r.Method == http.MethodGet {

		tmpl, err := template.ParseFiles("templates/form.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, nil)
		return
	}

	// POST REQUEST
	if r.Method == http.MethodPost {

		name := r.FormValue("name")
		ageValue := r.FormValue("age")
		email := r.FormValue("email")
		department := r.FormValue("department")
		language := r.FormValue("language")

		age, err := strconv.Atoi(ageValue)
		if err != nil {
			http.Error(w, "Invalid age", http.StatusBadRequest)
			return
		}
		if age < 18 {
			http.Error(w, "Under age limit", http.StatusConflict)
			return
		}
		student := Student{
			Name:             name,
			Age:              age,
			Email:            email,
			Department:       department,
			FavoriteLanguage: language,
		}

		tmpl, err := template.ParseFiles("templates/profile.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, student)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func main() {
	http.HandleFunc("/", handler)

	http.ListenAndServe(":8091", nil)
}
