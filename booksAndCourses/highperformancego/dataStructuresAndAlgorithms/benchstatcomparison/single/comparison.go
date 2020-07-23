package comparison

import (
	"encoding/json"
	"fmt"
	"log"
)

type CarData struct {
	cylinders int
	brand     string
	mpg       float64
}

func CarJSON() {
	honda := CarData{
		cylinders: 4,
		brand:     "honda",
		mpg:       42.6,
	}

	var CarDataJSON []byte
	CarDataJSON, err := json.Marshal((honda))
	if err != nil {
		log.Println(err)

	}
	fmt.Sprintf(string(CarDataJSON))
}
