package main

import (
	"html/template"
	"net/http"
)

type Student struct {
	Name       string
	Age        int
	IsAdmin    bool
	Courses    []string
	Department string
	Skills     []string
}

func handler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("templates/dashboard.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	student := Student{
		Name:       "Stephen",
		Age:        23,
		IsAdmin:    true,
		Department: "Computer Science",
		Courses: []string{
			"Go",
			"HTML",
			"Algorithms",
			"Databases",
			"Css",
			"Devops",
		},
		Skills: []string{
			"typing",
			"great with mechanics",
			"resiliant",
		},
	}

	err = tmpl.Execute(w, student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", handler)

	http.ListenAndServe(":8080", nil)
}
