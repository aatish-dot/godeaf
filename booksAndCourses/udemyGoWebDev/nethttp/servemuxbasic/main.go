package main

import (
	"io"
	"net/http"
)

type testHandler int
type testHandler1 int

func t(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "test tests and more tests")
}

func t1(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "test1 tests1 and more tests1")
}

func main() {

	http.HandleFunc("/test/", t)
	//alsoworks: http.Handle("/test/", http.HandlerFunc(t))

	http.HandleFunc("/test1/", t1)

	http.ListenAndServe(":8080", nil)

}
