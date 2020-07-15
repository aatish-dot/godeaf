package main

import "fmt"

//Stack allows to Push,Pop or Peek. The 0th index of a stack is the first in and last out.
//Essentially Stacks follow LIFO - last in first out.
type Stack struct {
	slice []int
}

//Push adds an item to the stack
func (s *Stack) Push(i int) {
	s.slice = append(s.slice, i)
}

//Pop removes and returns the top most item in the stack
func (s *Stack) Pop() int {
	p := s.slice[len(s.slice)-1]
	s.slice = s.slice[:len(s.slice)-1]
	return p
}

//Peek returns the top most item in the stack
func (s *Stack) Peek() int {
	return s.slice[len(s.slice)-1]
}

func (s *Stack) String() string {
	return fmt.Sprint(s.slice)
}

func main() {
	var s = new(Stack)
	s.Push(1)
	s.Push(2)
	fmt.Println(s.Peek())
	fmt.Println(s)
	s.Pop()
	fmt.Println(s)

}
