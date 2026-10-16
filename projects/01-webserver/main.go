package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "method not supported", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "\nhello!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "\nParseForm() error: %v", err)
		return
	}

	fmt.Fprintf(w, "POST request successfull!")
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")

	fmt.Fprintf(w, "\nName: %s", name)
	fmt.Fprintf(w, "\nEmail: %s", email)
	fmt.Fprintf(w, "\nPassword: %s", password)
}

func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Server is running at PORT: 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
