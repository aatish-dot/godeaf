package main

import (
	"errors"
	"fmt"
)

//Queue is a struct of int that helps us to define the type of the queue
type Queue struct {
	sl []int
}

func main() {
	var q = new(Queue)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	fmt.Println("q contains", q)
	_, err := q.Dequeue()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("after removing from queue", q)
	_, err = q.Dequeue()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("after removing from queue", q)

}

//Enqueue will allows us to add an element to the back of  Queue
func (q *Queue) Enqueue(i int) {
	//Add i to the Queue
	q.sl = append(q.sl, i)

}

//Dequeue removes the first item in the Queue or panics if there is isnt one
func (q *Queue) Dequeue() (int, error) {
	if (len(q.sl)) > 0 {
		i := q.sl[0]
		q.sl = q.sl[1:]
		return i, nil
	}
	return 0, errors.New("No elements in the queue")
}

func (q *Queue) String() string {
	return fmt.Sprint(q.sl)
}
