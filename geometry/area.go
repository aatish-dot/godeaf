package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
	sides  int
}

type circle struct {
	radius float64
}

type shape interface {
	area() float64
	perim() float64
}

func main() {
	c := circle{
		radius: 5,
	}

	s := square{
		length: 50,
		sides:  4,
	}

	dl := []shape{c, s}
	var m func(...shape)
	var n func(shape, func() float64) func() float64
	n = info

	m = func(d1 ...shape) {
		var printer func() float64
		var halvedarea func() float64
		for _, il := range d1 {
			printer = n(il, il.area)
			fmt.Println("perimeter is ", printer())
			halvedarea = halfarea(il)
			fmt.Println("halved area once is ", halvedarea())
			fmt.Println("halved area twice is ", halvedarea())

		}
	}
	m(dl...)

}

func (s square) area() float64 {
	return (math.Pow((s.length), 2))
}

func (c circle) area() float64 {
	return math.Pow(c.radius, 2) * (math.Pi)
}

func (s square) perim() float64 {
	return s.length * 4
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func info(s shape, f func() float64) func() float64 {
	fmt.Println("area is ", f())

	return func() float64 {
		return s.perim()
	}

}

func halfarea(s shape) func() float64 {
	areafunc := (s.area)
	area := areafunc()
	return func() float64 {
		area /= 2
		return area
	}

}
