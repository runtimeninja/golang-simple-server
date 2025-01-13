package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprintf(w, "POST request successful\n")
	name := r.FormValue("name")
	email := r.FormValue("email")
	comment := r.FormValue("comment")

	fmt.Fprintf(w, "Name: %s\n", name)
	fmt.Fprintf(w, "Email: %s\n", email)
	fmt.Fprintf(w, "Your Word: %s\n", comment)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
	}
	fmt.Fprintf(w, "thanks for making my day dude!🥺")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static-files/"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("starting server at port 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
