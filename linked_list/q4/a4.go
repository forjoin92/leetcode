package q4

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var head, tail *ListNode
	for l1 != nil || l2 != nil {
		var n int
		if l1 != nil && l2 == nil {
			n = l1.Val
			l1 = l1.Next
		}
		if l1 == nil && l2 != nil {
			n = l2.Val
			l2 = l2.Next
		}
		if l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				n = l1.Val
				l1 = l1.Next
			} else {
				n = l2.Val
				l2 = l2.Next
			}
		}
		if head == nil {
			tail = &ListNode{
				Val:  n,
				Next: nil,
			}
			head = tail
		} else {
			tail.Next = &ListNode{
				Val:  n,
				Next: nil,
			}
			tail = tail.Next
		}
	}
	return head
}
