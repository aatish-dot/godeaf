package main

import "fmt"

func main() {

	// & hacking

	// clearing last 4 digits to zero

	var x uint8 = 0xAC
	fmt.Printf("%b\n", x)
	x = x & 0xF0 // or x &= 0xF0
	fmt.Printf("%b\n", x)

	//printing odd numbers less than 25
	// the no is odd if its least bit is 1
	//apply the bitwise & operator with the integer value 1
	for y := 0; y < 25; y++ {
		if y&1 == 1 {
			fmt.Printf("%d\t%b\n", y, y)
		}

	}

	// | OR hacking
	// using | to set arbitary bits to 1

	z := 196 //11000100

	fmt.Printf("%b\n", z)

	// setting last 2 least significant bits to 3
	z |= 3
	fmt.Printf("%b\n", z)

}
