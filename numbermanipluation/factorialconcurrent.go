package main

import "fmt"

func main() {

	n := 4

	c := factorialC(n)

	for i := range c {
		fmt.Println(i)
	}

}

func factorialC(f int) chan int {
	ch := make(chan int)
	sum := 1

	go func() {
		for i := 1; i <= f; i++ {
			sum *= i
		}
		ch <- sum
		close(ch)
	}()
	return ch

}
