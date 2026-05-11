package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to my Go server")
	fmt.Println(r.Method)
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Server started on :8085")

	http.ListenAndServe(":8085", nil)
}
