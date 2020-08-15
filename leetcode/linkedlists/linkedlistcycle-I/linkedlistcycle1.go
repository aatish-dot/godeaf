package linkedlists

//ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	fast := head
	slow := head

	for {
		slow = slow.Next

		if fast.Next != nil {
			fast = fast.Next.Next
		} else {
			return false
		}

		if slow == nil || fast == nil {
			return false
		}

		if fast == slow {
			return true
		}

	}

}
