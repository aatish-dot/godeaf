package linkedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

//**
//  * Definition for singly-linked list.
//  * type ListNode struct {
//  *     Val int
//  *     Next *ListNode
//  * }
//  */
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow := head
	fast := head
	prev := head

	for {
		slow = slow.Next

		if fast.Next != nil {
			fast = fast.Next.Next
		} else {
			return nil
		}

		if slow == nil || fast == nil {
			return nil
		}

		if slow == fast {
			for slow != prev {
				slow, prev = slow.Next, prev.Next
			}
			return slow

		}

	}
}
