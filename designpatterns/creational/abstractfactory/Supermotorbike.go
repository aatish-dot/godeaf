package abstractfactory

type SuperMotorbike struct{}

func (c *SuperMotorbike) GetWheels() int {
	return 2
}

func (c *SuperMotorbike) GetSeats() int {
	return 2
}

func (c *SuperMotorbike) GetType() int {
	return SuperMotorbikeType
}
