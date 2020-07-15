package builder

import (
	"testing"
)

func TestBuilderPattern(t *testing.T) {
	manufacturingComplex := ManufactiringDirector{}
	carBuilder := &CarBuilder{}

	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()

	car := carBuilder.GetVehicle()

	if car.Wheels != 4 {
		t.Error("Wheels of car are not enough. Should have been 4 but got ", car.Wheels)
	}

	if car.Structure != "Car" {
		t.Error("Structure of car isnt right. Should have been Car but got ", car.Structure)

	}

	if car.Seats != 5 {
		t.Error("Seats of car are not enough. Should have been 5 but got ", car.Seats)

	}
}
