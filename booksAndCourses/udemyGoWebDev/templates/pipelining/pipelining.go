package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

var fm = template.FuncMap{
	"kitchenformat": MonthDateYear,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func main() {
	tNow := time.Now()
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", tNow)
	if err != nil {
		log.Fatal(err)
	}

}

func MonthDateYear(t time.Time) string {
	return t.Format(time.Kitchen)
}
