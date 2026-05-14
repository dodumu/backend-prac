package main

import (
	"html/template"
	"net/http"
	"strings"
)

type Student struct {
	Name           string
	Department     string
	IsAdmin        bool
	FavLanguages   []string
	Hobbies        []string
	HasScholarship bool
}

func handler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
		return
	}

	if r.Method == http.MethodPost {
		var data Student
		name := r.FormValue("name")
		department := r.FormValue("department")
		admin := r.FormValue("Admin")
		if admin != "" {
			data.IsAdmin = true
		}
		lang := r.FormValue("FavLanguages")
		langs := strings.Fields(lang)
		data.FavLanguages = langs
		hobby := r.FormValue("Hobbies")
		scholar := r.FormValue("scholar")
		if scholar != "" {
			data.HasScholarship = true
		}
		data.Hobbies = strings.Split(hobby, "\n")
		data.Name = name
		data.Department = department

		tmpl, err := template.ParseFiles("templates/dashboard.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8099", nil)
}
