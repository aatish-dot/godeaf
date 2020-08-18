package main

import (
	"io"
	"net/http"
)

// ListenAndServe on port ":8080" using the default ServeMux.

// Use HandleFunc to add the following routes to the default ServeMux:

// "/"
// "/dog/"
// "/me/

// Add a func for each of the routes.

// Have the "/me/" route print out your name.

func s(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This is the root")
}

func d(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This is the dog")
}

func m(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This should be my name")
}

func main() {

	http.HandleFunc("/", s)
	http.HandleFunc("/dog/", d)
	http.HandleFunc("/me/", m)
	http.ListenAndServe(":8080", nil)

}
