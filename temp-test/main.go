package main

import (
	"html/template"
	"net/http"
)

type Data struct {
	Name     string
	Email    string
	Password string
	Courses  []string
}

func baseHandler(w http.ResponseWriter, page string, user Data) {
	tmpl, err := template.ParseFiles(
		"templates/base.html",
		"templates/"+page,
		"templates/parts/nav.html",
		"templates/parts/footer.html",
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "base", user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		baseHandler(w, "home.html", Data{})
		return
	}

	if r.Method == http.MethodPost {

		email := r.FormValue("email")
		password := r.FormValue("password")

		if password != "David123" {
			http.Error(w, "invalid password", http.StatusBadRequest)
			return
		}

		http.Redirect(
			w,
			r,
			"/profile?email="+email,
			http.StatusSeeOther,
		)
		return
	}

	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {

	email := r.URL.Query().Get("email")

	user := Data{
		Name:  "David",
		Email: email,
		Courses: []string{
			"AI/ML",
			"DevOps",
			"CyberSecurity",
		},
	}

	baseHandler(w, "profile.html", user)
}

func main() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/profile", ProfileHandler)
	http.ListenAndServe(":8071", nil)
}
