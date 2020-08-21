package main

import (
	"io"
	"log"
	"net/http"
	"text/template"
)

// 1. Take the previous program in the previous folder and change it so that:
// * a template is parsed and served
// * you pass data into the template
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}
func s(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln("Error in parsing ", err)
	}

	tpl.ExecuteTemplate(w, "index.gohtml", r.Form)

}

func d(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This is the dog")
}

func m(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This should be my name")
}

func main() {

	http.Handle("/", http.HandlerFunc(s))
	http.Handle("/dog/", http.HandlerFunc(s))
	http.Handle("/me/", http.HandlerFunc(m))
	http.ListenAndServe(":8080", nil)

}
