package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	// parsed the form data from the request.
	if err := r.ParseForm(); err != nil {
		//if there's an error parsing the form, send an error response
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

//helloHandler handles the "/hello" endpoint for GET requests
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
	}
	// Responded with a custom message for valid GET requests to "/hello".
	fmt.Fprintf(w, "thanks for making my day dude!🥺")
}

func main() {
	//Serve static files from the "static-files" directory at the root ("/")
	fileServer := http.FileServer(http.Dir("./static-files/"))
	http.Handle("/", fileServer)

	// Registered the formHandler for the "/form" endpoint.
	http.HandleFunc("/form", formHandler)
	//Registered the helloHandler for the "/hello" endpoint
	http.HandleFunc("/hello", helloHandler)

	//started the server on port 8080 and log any errors.
	fmt.Printf("starting server at port 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
