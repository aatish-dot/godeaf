package main

import (
	"log"
	"os"
	"text/template"
)

//contains information about California hotels including Name, Address, City, Zip, Region
//region can be: Southern, Central, Northern
//can hold an unlimited number of hotels

type hotel struct {
	Name    string
	Address string
	City    string
	Zip     int
	Region  string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").ParseFiles("tpl.gohtml"))
}

func main() {

	hotels := []hotel{
		hotel{
			Name:    "mariott",
			Address: "downtown avenue, in the middle",
			City:    "SFO",
			Zip:     11111,
			Region:  "Southern",
		},
		hotel{
			Name:    "westin",
			Address: "not so downtown avenue, in the corner",
			City:    "LA",
			Zip:     11112,
			Region:  "Northern",
		},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", hotels)

	if err != nil {
		log.Fatal(err)
	}

}
