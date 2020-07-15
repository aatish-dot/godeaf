package composition

import "testing"

func TestAthlete_Train(t *testing.T) {
	localSwim := Swim
	swimmer := CompositeSwimmerA{
		MySwim: &localSwim,
	}
	swimmer.MyAthelete.Train()
	(*swimmer.MySwim)()
}

func TestAnimal_Swim(t *testing.T) {
	fish := Shark{
		Swim: Swim,
	}
	fish.Eat()
	fish.Swim()
}

func TestSwimmer_Swim2(t *testing.T) {
	swimmer := CompositeSwimmerB{
		&Athelete{},
		&SwimmerImplementor{},
	}
	swimmer.Train()
	swimmer.Swim()
}
