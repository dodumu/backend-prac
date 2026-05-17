package main

import (
	"html/template"
	"net/http"
)

type User struct {
	Name string
}

func BaseHandler(w http.ResponseWriter, page string, data User) {
	tmpl, err := template.ParseFiles(
		"templates/base.html",
		"templates/parts/footer.html",
		"templates/parts/navbar.html",
		"templates/"+page,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	name := "David"
	data := User{
		Name: name,
	}

	if r.Method == http.MethodPost {
		http.Redirect(w, r, "/courses?name="+name, http.StatusSeeOther)
	}
	BaseHandler(w, "home.html", data)
}

func courseHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	data := User{
		Name: name,
	}

	BaseHandler(w, "courses.html", data)
}

func main() {
	fs := http.FileServer(http.Dir("static"))

	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/courses", courseHandler)

	http.ListenAndServe(":8077", nil)
}
