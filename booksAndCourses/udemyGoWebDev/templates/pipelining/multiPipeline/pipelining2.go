package main

import (
	"log"
	"math"
	"os"
	"text/template"
)

func sqr(n float64) float64 {
	n = math.Pow(n, 2)
	return n
}

func double(n float64) float64 {
	return n * 2
}

func sqrt(n float64) float64 {
	return math.Sqrt(n)
}

var fm2 = template.FuncMap{
	"double":     double,
	"square":     sqr,
	"squareroot": sqrt,
}

var tpl2 *template.Template

func init() {
	tpl2 = template.Must(template.New("").Funcs(fm2).ParseFiles("tpl2.gohtml"))
}

func main() {
	err := tpl2.ExecuteTemplate(os.Stdout, "tpl2.gohtml", 3.0)
	if err != nil {
		log.Fatal(err)
	}
}
