package iterators

import "fmt"

//benchmarking different iterators in golang
var sumloops int
var sumcallback int
var sumNext int
var sumBufferedChan int

//simple for loop for adding all numbers form 0 to n
func simpleforloop(n int) int {
	for i := 0; i < n; i++ {
		sumloops += i
	}
	return sumloops
}

//Callbackloop uses a call back loop to add numbers from 0 to n
func CallbackLoop(top int) {
	err := callbackloopiterator(top, func(n int) error {
		sumcallback += n
		return nil
	})
	if err != nil {
		panic(err)
	}
}

func callbackloopiterator(top int, callback func(n int) error) error {
	for i := 0; i < top; i++ {
		err := callback(i)
		if err != nil {
			return err
		}
	}
	return nil

}

type Counterstruct struct {
	err error
	max int
	cur int
}

func NewCounterIterator(top int) *Counterstruct {
	var err error
	return &Counterstruct{
		err: err,
		max: top,
		cur: 0,
	}
}

func (i *Counterstruct) Next() bool {
	if i.err != nil {
		return false
	}
	i.cur++
	return i.cur <= i.max
}

func (i *Counterstruct) Value() int {
	if i.err != nil || i.cur > i.max {
		panic("Value is not valid after iterator finished")
	}
	return i.cur
}

func NextLoop(top int) {
	nextIterator := NewCounterIterator(top)
	for nextIterator.Next() {
		fmt.Print(nextIterator.Value())
	}
}

//use channels for the same job
func BufferedChanLoop(n int) int {
	ch := make(chan int, n)
	go func() {
		defer close(ch)
		for i := 0; i < n; i++ {
			ch <- i
		}
	}()

	for j := range ch {
		sumBufferedChan += j
	}
	return sumBufferedChan
}
