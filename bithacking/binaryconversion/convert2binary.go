package main

import (
	"fmt"
	"strconv"
)

var initv string = "0"
var bin string

func main() {
	var int2convert int
	fmt.Println("Please enter the interger to convert to binary:")
	fmt.Scan(&int2convert)

	//binary value as converted from the recursove function
	fmt.Println(convert2bin(int2convert))

	//Binary value as got from %b
	fmt.Printf("%b", int2convert)
}

func convert2bin(n int) string {

	if n == 0 || n == 1 {
		bin = strconv.Itoa(n)
	} else {

		k := n / 2
		b := strconv.Itoa(n % 2)

		bin = convert2bin(k) + b

	}

	fmt.Println(n, bin)
	return bin

}
