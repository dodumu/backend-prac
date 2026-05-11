package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {

	pass := "49321111@David"

	fmt.Println(r.Method)

	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")
		if password != pass {
			w.Write([]byte("invalid password"))
			return
		}
		age := r.FormValue("age")
		agee, err := strconv.Atoi(age)
		if err != nil {
			w.Write([]byte("invalid age"))
			return
		}
		if agee < 18 {
			w.Write([]byte("Under Age"))
			return
		}
		color := r.FormValue("color")
		adress := r.FormValue("text")
		w.Write([]byte("Hello " + username + "\n"))
		w.Write([]byte("Your email is: " + email + "\n"))
		w.Write([]byte("You are: " + age + " years old\n"))
		w.Write([]byte("You live at: " + adress + "\n"))
		w.Write([]byte("Your favorite color is: " + color))
		return
	}

	http.ServeFile(w, r, "index.html")

}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Server started on :8085")

	http.ListenAndServe(":8085", nil)
}
