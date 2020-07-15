package builder

//BuildProcess type is something that will define what all object types need to implement in terms of functions
type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

//ManufactiringDirector helps the user to abstract the build process from the consumer of the final types in this case car / bike.
type ManufactiringDirector struct {
	builder BuildProcess
}

//Construct method helps in the process of the creation of objects
func (f *ManufactiringDirector) Construct() {
	f.builder.SetSeats().SetStructure().SetWheels()
}

//SetBuilder defines the builder
func (f *ManufactiringDirector) SetBuilder(b BuildProcess) {
	f.builder = b
}

//VehicleProduct defines the structure for all vehicles.
type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

//CarBuilder of type car
//Note that more attributes can be defined which need not be for all vehicles.
type CarBuilder struct {
	v VehicleProduct
}

//SetWheels returns 4 for a car
func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
	return c
}

//SetSeats returns 5 seats for a car
func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.Seats = 5
	return c
}

//SetStructure helps the user to define how a car is.
func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Car"
	return c
}

//GetVehicle returns an instance of the Car
func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.v
}

//BikeBuilder of type motorbike
type BikeBuilder struct {
	v VehicleProduct
}

//SetWheels helps the user to define how many wheels a bike has.
func (c *BikeBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 2
	return c
}

//SetSeats defines that  a Bike has 2 seats
func (c *BikeBuilder) SetSeats() BuildProcess {
	c.v.Seats = 2
	return c
}

//SetStructure will define the object to be of type Bike.
func (c *BikeBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Motorbike"
	return c
}

//GetVehicle returns an instance of the Bike
func (c *BikeBuilder) GetVehicle() VehicleProduct {
	return c.v
}
