package main

import "fmt"

func main() {

	unsortedint := []int{3, 1, 5, 9, 2, 4, 10}
	fmt.Println(unsortedint)
	sortedint := insertionsort(unsortedint) //note that the unsortedint slice also changes
	fmt.Println(sortedint)

}

func insertionsort(items []int) []int {
	for i := 1; i < len(items); i++ {
		j := i
		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j--
		}
	}
	return items
}
