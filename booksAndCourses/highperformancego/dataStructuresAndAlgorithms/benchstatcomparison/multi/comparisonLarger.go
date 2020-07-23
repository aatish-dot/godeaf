package comparisonLarger

import (
	"encoding/json"
	"fmt"
	"log"
)

type MoreCarData struct {
	cylinders int
	origin    string
	mpg       float64
	doors     int
	fueltype  string
}

func CarJSON() {
	sedan := MoreCarData{cylinders: 2, origin: "japanese", mpg: 45.2, doors: 4, fueltype: "petrol"}
	convertible := MoreCarData{cylinders: 4, origin: "american", mpg: 67.8, doors: 2, fueltype: "diesel"}
	cars := []MoreCarData{sedan, convertible}
	for _, car := range cars {
		var CarDataJSON []byte
		CarDataJSON, err := json.Marshal(car)
		if err != nil {
			log.Println(err)
		}
		fmt.Sprintf(string(CarDataJSON))
	}
}
