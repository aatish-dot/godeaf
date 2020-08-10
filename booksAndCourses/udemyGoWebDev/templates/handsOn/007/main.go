package main

import (
	"log"
	"os"
	"text/template"
)

//Create a data structure to pass to a template which
//contains information about restaurant's menu including Breakfast, Lunch, and Dinner items

type item struct {
	Name, Description, Meal string
	Price                   float64
}

type restaurant struct {
	Name string
	Menu []item
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	menu1 := []item{
		item{
			Name:        "Boiled Eggs",
			Description: "Overpriced Eggs boiled and overboiled",
			Meal:        "Breakfast",
			Price:       45.2,
		},
		item{
			Name:        "Boiled Chicken",
			Description: "cheap high protein meal",
			Meal:        "Lunch",
			Price:       1.5,
		},
	}

	menu2 := []item{
		item{
			Name:        "Hard Boiled Eggs",
			Description: "Hard Eggs boiled and overboiled",
			Meal:        "Breakfast",
			Price:       45.5,
		},
		item{
			Name:        "Fried Chicken",
			Description: "cheap keto meal",
			Meal:        "Lunch",
			Price:       1.5,
		},
	}

	SFORestaurants := []restaurant{
		restaurant{
			Name: "Subway",
			Menu: menu1,
		},
		restaurant{
			Name: "KFC",
			Menu: menu2,
		},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", SFORestaurants)
	if err != nil {
		log.Fatal(err)
	}
}
