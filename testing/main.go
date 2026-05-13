package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	if r.Method == http.MethodPost {
		name := r.FormValue("full-name")
		email := r.FormValue("email")
		password := r.FormValue("password")
		age := r.FormValue("age")
		gender := r.FormValue("gender")
		department := r.FormValue("department")
		address := r.FormValue("address")
		TC := r.FormValue("T&C")
		years, err := strconv.Atoi(age)
		if err != nil {
			fmt.Fprintln(w, "invalid age")
			return
		}
		up := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		low := "abcdefghiklmnopqrstuvwxyz"
		num := "0123456789"
		if !strings.ContainsAny(password, up) || !strings.ContainsAny(password, low) || !strings.ContainsAny(password, num) {
			fmt.Fprintln(w, "invalid password format")
			return
		}
		if years < 18 || password != "48321111David" {
			fmt.Fprintln(w, "invalid crendtials")
			return
		}
		if TC == "" {
			fmt.Fprintln(w, "you need to accept terms and conditions")
			return
		}
		fmt.Fprintf(w, "welcome to the portal %v\n", name)
		fmt.Fprintf(w, "your email is %v\n", email)
		fmt.Fprintf(w, "you are %v years old\n", years)
		fmt.Fprintln(w, gender)
		fmt.Fprintf(w, "you are a %v student\n", department)
		fmt.Fprintf(w, "your address is %v\n", address)
		fmt.Fprintf(w, "do you agree to out terms and conditons?\n")
		fmt.Fprintln(w, TC)
		return
	}
	http.ServeFile(w, r, "index.html")
}

func main() {
	fmt.Println("The program is running on port: 8087:")
	http.HandleFunc("/", handler)

	http.ListenAndServe(":8087", nil)
}
