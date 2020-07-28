package main

import "fmt"

func main() {

	unsortedint := []int{3, 1, 5, 9, 2, 4, 10}
	fmt.Println(unsortedint)
	quicksort(unsortedint, 0, len(unsortedint)-1) //note that the unsortedint slice also changes
	fmt.Println(unsortedint)

}

func quicksort(items []int, start, end int) {
	fmt.Println("start", start, "end", end)
	if (end - start) < 1 {

		return

	}

	pivot := items[end]
	splitIndex := start

	for i := start; i < end; i++ {
		fmt.Println(splitIndex, i)
		if items[i] < pivot {
			items[i], items[splitIndex] = items[splitIndex], items[i]
			splitIndex++
			fmt.Println(items)
		}
	}

	items[end] = items[splitIndex]
	items[splitIndex] = pivot
	fmt.Println(items)

	quicksort(items, start, splitIndex-1)
	quicksort(items, splitIndex+1, end)

}
