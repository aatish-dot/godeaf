package main

import (
	"fmt"
	"time"
)

type Ball struct {
	hits int
}

func main() {
	table := make(chan *Ball)

	go player("ping", table)
	go player("pong", table)

	table <- new(Ball)          //game on; toss the Ball
	time.Sleep(1 * time.Second) //game over; grab the ball
	<-table

}

func player(name string, table chan *Ball) {
	fmt.Println("in player")
	for {
		ball := <-table
		ball.hits++
		fmt.Println(name, ball.hits)
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}
