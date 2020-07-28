package main

import "fmt"

func main() {

	unsortedint := []int{3, 1, 5, 9, 2, 4, 10}
	fmt.Println(unsortedint)
	sortedint := mergesort(unsortedint) //note that the unsortedint slice also changes
	fmt.Println(sortedint)

}

func mergesort(items []int) []int {
	length := len(items)
	if length == 1 {
		return items
	}
	middle := length / 2

	left := make([]int, middle)
	right := make([]int, (length - middle))

	for i := 0; i < length; i++ {
		if i < middle {
			left[i] = items[i]
		} else {
			right[i-middle] = items[i]
		}
	}

	return merge(mergesort(left), mergesort(right))

}

func merge(left, right []int) (result []int) {
	result = make([]int, len(left)+len(right))
	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}

	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	fmt.Println(result)
	return
}
