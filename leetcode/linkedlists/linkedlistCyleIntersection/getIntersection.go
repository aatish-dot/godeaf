package linkedlist

//* ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a := headA
	b := headB
	maxa, maxb := 0, 0

	//Get length of A
	for a != nil {
		a = a.Next
		maxa++

	}

	//Get length of B
	for b != nil {
		b = b.Next
		maxb++
	}

	a = headA
	b = headB

	//Check which is the bigger list and reduce the size to an equal length as beyond the intersection the length will be same
	if maxa > maxb {
		diff := maxa - maxb
		for diff != 0 {
			diff--
			a = a.Next
		}

	} else {
		diff := maxb - maxa
		for diff != 0 {
			diff--
			b = b.Next
		}

	}

	//till the intersection or till the end of the lists, continue to check if the intersection point exists (a=b) else return nil
	for a != nil || b != nil {

		if a == b {
			return a
		}

		a = a.Next
		b = b.Next
	}

	return nil

}
