package main

import "fmt"

func main() {
	m := 119
	n := 544

	gcd := getGCD(m, n)

	fmt.Printf("GCD of %d and %d is %d", m, n, gcd)
}

func getGCD(m, n int) int {
	if n == 0 {
		return 0
	}
	r := m % n
	if r == 0 {
		return n
	}
	return getGCD(n, r)
}
