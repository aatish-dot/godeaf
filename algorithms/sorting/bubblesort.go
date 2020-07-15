package main

import (
	"fmt"
	"strings"
)

func main() {
	numbers := []int{2, 1, 3, 4, 5, 6}
	sortednumbers := bubblesortInt(numbers)
	fmt.Println("New list is ", sortednumbers)

	var animals []string = []string{
		"dog",
		"cat",
		"alligator",
		"cheetah",
		"rat",
		"moose",
		"cow",
		"mink",
		"porcupine",
		"dung beetle",
		"camel",
		"steer",
		"bat",
		"hamster",
		"horse",
		"colt",
		"bald eagle",
		"frog",
		"rooster",
	}

	sortedanimals := bubblesortString(animals)
	fmt.Println("New list is ", sortedanimals)

}

func bubblesortInt(numbers []int) []int {
	for i := 0; i <= len(numbers)-2; i++ {
		alreadysortedindex := (len(numbers) - i)
		sortedlist, sorted := sweepInt(numbers[:alreadysortedindex])
		sortedlist = append(sortedlist, numbers[alreadysortedindex:]...)
		if !sorted {
			break
		}
	}
	return numbers
}

func bubblesortString(str []string) []string {
	for i := 0; i <= len(str)-2; i++ {
		alreadysortedindex := (len(str) - i)
		sortedlist, sorted := sweepString(str[:alreadysortedindex])
		sortedlist = append(sortedlist, str[alreadysortedindex:]...)
		if !sorted {
			break
		}
	}
	return str
}

func sweepInt(s []int) ([]int, bool) {
	swapped := false
	fmt.Println("sorting: ", s)

	for i := 0; i < (len(s) - 1); i++ {
		if s[i] > s[i+1] {
			s[i], s[i+1] = s[i+1], s[i]
			swapped = true
		}
	}
	return s, swapped
}

func sweepString(s []string) ([]string, bool) {
	swapped := false
	fmt.Println("sorting: ", s)

	for i := 0; i < (len(s) - 1); i++ {
		if stringGreaterThan(s[i], s[i+1]) {
			s[i], s[i+1] = s[i+1], s[i]
			swapped = true
		}
	}
	return s, swapped
}

func stringGreaterThan(s1, s2 string) bool {
	if strings.ToLower(s1) > strings.ToLower(s2) {
		return true
	}
	return false

}
