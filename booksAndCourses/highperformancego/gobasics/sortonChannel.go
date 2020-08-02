package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{2, 3, 4, 33, 55, 66, 24, 22, 67, 99}
	fmt.Println("unsorted slice: ", s)
	chanbool := make(chan bool)
	inttosearch := 24
	go func() {
		sortInts(s, chanbool)
		searchInts(s, inttosearch, chanbool)
	}()
	<-chanbool
}

func sortInts(intslice []int, done chan bool) {
	sort.Ints(intslice)
	fmt.Printf("sorted slice : %v\n", intslice)
	done <- true
}

func searchInts(intslice []int, searchint int, done chan bool) {
	searched := sort.SearchInts(intslice, searchint)
	if searched < len(intslice) {
		fmt.Printf("Found %d at array position %d\n", searchint, searched)
	} else {
		fmt.Printf("Could not find %d at any position", searchint)
	}
	done <- true
}
