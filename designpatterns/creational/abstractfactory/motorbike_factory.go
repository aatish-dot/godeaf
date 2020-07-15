package abstractfactory

import (
	"fmt"
)

const (
	SuperMotorbikeType  = 1
	CruiseMotorbikeType = 2
)

type MotorbikeFactory struct{}

func (c *MotorbikeFactory) GetVehicle(v int) (Vehicle, error) {
	switch v {
	case SuperMotorbikeType:
		return new(SuperMotorbike), nil
	case CruiseMotorbikeType:
		return new(CruiseMotorbike), nil
	default:
		return nil, fmt.Errorf("Vehicle of type %d is not recognized", v)
	}
}
