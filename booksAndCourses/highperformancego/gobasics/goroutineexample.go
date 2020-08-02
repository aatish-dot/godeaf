package main

import (
	"fmt"
	"time"
)

func main() {
	const sleeptime time.Duration = 14
	go printSleep("HELLO GOPHERS")
	time.Sleep(sleeptime * time.Millisecond)
	fmt.Println("sleep complete")

}

func printSleep(s string) {
	for index, value := range s {
		fmt.Printf("%#U is at index %d\n ", value, index)
		time.Sleep(1 * time.Millisecond)
	}
}
