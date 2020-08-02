package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan bool)
	ch3 := make(chan rune)

	go func() {
		ch1 <- "Hello gophers"
	}()

	go func() {
		ch2 <- true
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch3 <- "r"
	}()

	for i := 0; i < 3; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("String channel returned: ", msg1)
		case msg2 := <-ch2:
			fmt.Println("Bool channel returned: ", msg2)
		case msg3 := <-ch3:
			fmt.Println("Rune channel returned ", msg3)
		}
	}

}
