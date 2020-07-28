package main

import "fmt"

//binary search has a complexity of O(log n) compared to O(n) for linear search
// to use the algo the list needs to be sorted

func main() {
	items := []int{1, 2, 3, 4, 5, 6, 7, 7, 7, 9}
	lmatch := lb_binarysearch(7, items, 0, len(items)-1)

	fmt.Println("-----------------")
	fmt.Println("lowest matching index is ", lmatch, "\n highest matching index is ", hb_binarysearch(7, items, lmatch, len(items)-1))
}

func lb_binarysearch(needle int, haystack []int, low, high int) int {

	if low > high {
		return low
	}

	median := (low + high) / 2

	fmt.Println(median, high, low)

	if haystack[median] >= needle {

		return lb_binarysearch(needle, haystack, low, median-1)
	} else {
		return lb_binarysearch(needle, haystack, median+1, high)
	}

}

//we need to find out multiple occurances and hence pass the lowest occurrence as "low"
//helps to keep the search still at O(log n)
func hb_binarysearch(needle int, haystack []int, low, high int) int {

	fmt.Println(low)
	if low > high {
		if low == len(haystack)-1 && haystack[low] == needle {
			return low
		}
		return low - 1
	}

	median := (low + (high-low)/2)

	fmt.Println(median, high, low)

	if haystack[median] > needle {

		return hb_binarysearch(needle, haystack, low, median-1)
	} else {
		return hb_binarysearch(needle, haystack, median+1, high)
	}

}
