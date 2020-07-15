package composition

import "fmt"

type Athelete struct{}

func (a *Athelete) Train() {
	fmt.Println("Training")
}

func Swim() {
	fmt.Println("Swimming")
}

type CompositeSwimmerA struct {
	MyAthelete Athelete
	MySwim     *func()
}

//--------------------------------------------

type Trainer interface {
	Train()
}

type Swimmer interface {
	Swim()
}

type SwimmerImplementor struct{}

func (s *SwimmerImplementor) Swim() {
	println("Swimming")
}

type CompositeSwimmerB struct {
	Trainer
	Swimmer
}

//------------------------------------------------

type Animal struct{}

func (r *Animal) Eat() {
	fmt.Println("Eating")
}

type Shark struct {
	Animal
	Swim func()
}

//----------------------------------------------------
