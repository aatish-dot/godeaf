package main

import "fmt"

func main() {

	disks := 3
	start := "A"
	end := "C"
	buffer := "B"
	recurssionhanoi(disks, start, end, buffer, printvalues)

}

func recurssionhanoi(n int, start string, end string, buffer string, callback func(int, string, string)) {

	if n == 1 {
		callback(n, start, end)
	} else {
		recurssionhanoi(n-1, start, buffer, end, callback)
		callback(n, start, end)
		recurssionhanoi(n-1, buffer, end, start, callback)
	}

}

func printvalues(n int, movedfrom string, movedto string) {
	fmt.Printf("Moved disk %v from %s to %s\n", n, movedfrom, movedto)
}
