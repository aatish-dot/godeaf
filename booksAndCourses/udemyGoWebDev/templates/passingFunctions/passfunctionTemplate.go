package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name string
	Moto string
}

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	if len(s) > 3 {
		s = s[:3]
	}
	return s
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))

}
func main() {

	b := sage{
		Name: "Buddha",
		Moto: "Let there be peace !",
	}

	g := sage{
		Name: "Gandhi",
		Moto: "Let there be more peace !",
	}

	sages := []sage{b, g}

	data := struct {
		Wisdom []sage
	}{
		sages,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		fmt.Println(err)
	}

}
