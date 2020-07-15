package main

import "fmt"

func main() {
	x := []int{0, 1, 2, 3, 3, 4, 5, 6, 7, 8}

	s := even(sum, x...)

	fmt.Println("some of even numbers", s)

	p := odd(sum, x...)

	fmt.Println("some of odd numbers", p)

}

func even(f func(x ...int) int, y ...int) int {
	var z []int
	for _, v := range y {
		if v%2 == 0 {
			z = append(z, v)
		}
	}
	return f(z...)

}

func sum(u ...int) int {
	total := 0
	for _, v := range u {
		total += v

	}
	return total
}

func odd(f func(x ...int) int, y ...int) int {

	var od []int
	for _, s := range y {
		if s%2 != 0 && s != 0 {
			od = append(od, s)
		}

	}
	return f(od...)
}
