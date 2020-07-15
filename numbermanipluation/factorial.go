package main

import "fmt"

func main() {

	n := factorial(5)
	fmt.Printf("factorial for 5 is %d \n ", n)
	nloop := loopfac(5)
	fmt.Printf("loop factorial for 5 is %d ", nloop)

}

func factorial(i int) int {
	if i == 0 {
		return 1
	}
	return i * factorial(i-1)

}

func loopfac(i int) int {
	fac := 1
	for ; i > 0; i-- {
		fac = fac * i
	}
	return fac
}
