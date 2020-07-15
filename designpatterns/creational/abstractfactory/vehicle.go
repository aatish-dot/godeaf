package abstractfactory

//Vehicle interface to define Vehicle as something that implements GetWheels and GetSeats
type Vehicle interface {
	GetWheels() int
	GetSeats() int
}
