package main

import "fmt"

func main() {

	unsortedint := []int{3, 1, 5, 9, 2, 4, 10}
	fmt.Println(unsortedint)
	sortedint := heapsort(unsortedint) //note that the unsortedint slice also changes
	fmt.Println(sortedint)

}

func heapsort(items []int) []int {
	n := len(items)

	//build a max heap
	for i := n/2 - 1; i >= 0; i-- {
		heapify(items, n, i)
	}

	//one by by one extract items
	for i := n - 1; i > 0; i-- {
		//move current root to end
		temp := items[0]
		items[0] = items[i]
		items[i] = temp

		heapify(items, i, 0)
	}
	return items
}

//to heapify subtree rooted at index i
//#n is size of tree
func heapify(items []int, n, i int) {
	largest := i //intialize largest as target root
	left := 2*i + 1
	right := 2*i + 2

	if left < n && items[i] < items[left] {
		largest = left
	}

	if right < n && items[largest] < items[right] {
		largest = right
	}
	//change root if needed
	if largest != i {
		items[i], items[largest] = items[largest], items[i]
		//heapify the affected sub tree
		heapify(items, n, largest)
	}

}
