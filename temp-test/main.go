package main

import (
	"html/template"
	"net/http"
)

type User struct {
	Name     string
	Courses  []string
	Password string
	Email    string
}

func baseTemp(w http.ResponseWriter, file string, data User) {
	tmpl, err := template.ParseFiles(
		"templates/base.html",
		"templates/"+file,
		"templates/parts/foot.html",
		"templates/parts/nav.html",
	)
	if err != nil {
		http.Error(w, "could not load file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func parseForm(r *http.Request) (User, bool) {
	var user User
	user.Name = r.FormValue("name")
	user.Email = r.FormValue("email")
	user.Password = r.FormValue("password")
	verify := r.FormValue("verifypass")

	if verify != user.Password {
		return user, false
	}

	courses := []string{"CyberSecurity", "AI/ML", "DevOps", "Data-Science", "Crypto-Currency", "Video-Game", "Mobile-App", "Full-Stack"}
	for _, c := range courses {
		if r.FormValue(c) != "" {
			user.Courses = append(user.Courses, c)
		}
	}
	return user, true
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	if r.Method == http.MethodPost {
		var ok bool
		user, ok = parseForm(r)
		if !ok {
			http.Error(w, "passwords do not match", http.StatusBadRequest)
			return
		}
	}
	baseTemp(w, "home.html", user)
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	if r.Method == http.MethodPost {
		var ok bool
		user, ok = parseForm(r)
		if !ok {
			http.Error(w, "passwords do not match", http.StatusBadRequest)
			return
		}
	}
	baseTemp(w, "profile.html", user)
}

func courseHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	if r.Method == http.MethodPost {
		var ok bool
		user, ok = parseForm(r)
		if !ok {
			http.Error(w, "passwords do not match", http.StatusBadRequest)
			return
		}
	}
	baseTemp(w, "courses.html", user)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/profile", profileHandler)
	http.HandleFunc("/courses", courseHandler)

	http.ListenAndServe(":8071", nil)
}
