package main

import (
	"fmt"
)

type MyLinkedList struct {
	Head *Node
	Tail *Node
	size int
}

type Node struct {
	val  int
	next *Node
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	cur := this.Head

	if index == 0 {
		return cur.val
	}
	if index > this.size-1 || index < 0 {
		return -1
	}
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	return cur.val

}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(value int) {
	n := &Node{
		val: value,
	}
	if this.Head == nil {
		this.Head = n
	} else if this.size == 1 {
		cur := this.Head
		this.Head = n
		this.Head.next = cur
		this.Tail = cur
	} else {
		cur := this.Head
		this.Head = n
		this.Head.next = cur
	}
	this.size++
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(value int) {
	n := &Node{
		val: value,
	}
	if this.Tail == nil {
		this.Tail = n
		this.Head.next = n
	} else {
		cur := this.Tail
		cur.next = n
		this.Tail = n

	}
	this.size++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, value int) {
	var cur *Node
	cur = this.Head
	n := &Node{
		val: value,
	}

	if index > this.size || index < 0 {
		return
	} else if index == 0 || this.Head == nil {
		this.AddAtHead(value)
		return
	} else if index == this.size {
		this.AddAtTail(value)
		return
	} else {
		for i := 0; i < index-1; i++ {
			cur = cur.next
		}
		cur1 := cur.next
		n.next = cur1
		cur.next = n
	}
	this.size++

}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	var cur, oldcur *Node
	cur = this.Head
	if index > this.size-1 {
		return
	} else if index == 0 {
		this.Head = cur.next

	} else {
		for i := 0; i < index; i++ {
			oldcur = cur
			cur = cur.next
		}
		oldcur.next = cur.next
	}
	if this.size > 0 {
		this.size--
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

func main() {
	linkedList := Constructor() // Initialize empty LinkedList

	linkedList.AddAtHead(84)
	linkedList.printlinkedListMembers()
	linkedList.AddAtTail(2)
	linkedList.printlinkedListMembers()
	linkedList.AddAtTail(39)
	linkedList.printlinkedListMembers()
	linkedList.Get(3)
	linkedList.printlinkedListMembers()
	linkedList.Get(1)
	linkedList.printlinkedListMembers()
	linkedList.AddAtTail(42)
	linkedList.printlinkedListMembers()
	linkedList.AddAtIndex(1, 80)
	linkedList.printlinkedListMembers()
	linkedList.AddAtHead(14)
	linkedList.printlinkedListMembers()
	linkedList.AddAtHead(1)
	linkedList.printlinkedListMembers()
	linkedList.AddAtTail(53)
	linkedList.printlinkedListMembers()
	linkedList.AddAtTail(98)
	linkedList.printlinkedListMembers()
	linkedList.AddAtTail(19)
	linkedList.printlinkedListMembers()
	linkedList.AddAtTail(12)
	linkedList.printlinkedListMembers()
	linkedList.Get(2)
	linkedList.printlinkedListMembers()
	linkedList.AddAtHead(16)
	linkedList.printlinkedListMembers()
	linkedList.AddAtHead(33)
	linkedList.printlinkedListMembers()
	linkedList.AddAtIndex(4, 17)
	linkedList.printlinkedListMembers()
	linkedList.AddAtIndex(6, 8)
	linkedList.printlinkedListMembers()
	linkedList.AddAtHead(37)
	linkedList.printlinkedListMembers()
	linkedList.AddAtTail(43)
	linkedList.printlinkedListMembers()
	linkedList.DeleteAtIndex(11)
	linkedList.printlinkedListMembers()
	linkedList.AddAtHead(80)
	linkedList.printlinkedListMembers()
	linkedList.AddAtHead(31)
	linkedList.printlinkedListMembers()
	linkedList.AddAtIndex(13, 23)
	linkedList.printlinkedListMembers()
	linkedList.AddAtTail(17)
	linkedList.printlinkedListMembers()
	linkedList.Get(4)
	linkedList.printlinkedListMembers()
	linkedList.AddAtIndex(10, 0)
	linkedList.printlinkedListMembers()
	linkedList.AddAtTail(21)
	linkedList.printlinkedListMembers()
	linkedList.AddAtHead(73)
	linkedList.printlinkedListMembers()
	linkedList.AddAtHead(22)
	linkedList.printlinkedListMembers()
	linkedList.AddAtIndex(24, 37)
	linkedList.printlinkedListMembers()
	linkedList.AddAtTail(14)
	linkedList.printlinkedListMembers()
	linkedList.AddAtHead(97)
	linkedList.printlinkedListMembers()
	linkedList.AddAtHead(8)
	linkedList.printlinkedListMembers()
	linkedList.Get(6)
	linkedList.printlinkedListMembers()
	linkedList.DeleteAtIndex(17)
	linkedList.printlinkedListMembers()
}

func (ll MyLinkedList) printlinkedListMembers() {
	var listmem []int
	cur := ll.Head
	for i := 0; i < ll.size; i++ {
		listmem = append(listmem, cur.val)
		cur = cur.next
	}
	fmt.Println(listmem, ll.size, ll.Head.val, ll.Tail)
}
