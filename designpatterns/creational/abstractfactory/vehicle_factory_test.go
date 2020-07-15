package abstractfactory

import (
	"testing"
)

func TestMotorbikeFactory(t *testing.T) {
	motorbikeF, err := GetVehicleFactory(MotorbikeFactoryType)
	if err != nil {
		t.Fatal(err)
	}

	motorbikeVehicle, err := motorbikeF.GetVehicle(SuperMotorbikeType)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Motorbike has %d wheels and %d seats\n", motorbikeVehicle.GetWheels(), motorbikeVehicle.GetSeats())

	superBike, ok := motorbikeVehicle.(Motorbike)
	if !ok {
		t.Fatal("Struct assertion has failed\n")
	}
	t.Logf("Super motorbike has type %d\n", superBike.GetType())
}
