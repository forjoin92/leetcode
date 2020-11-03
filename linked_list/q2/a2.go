package q2

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p1, p2 := head, head
	for i := 0; i < n-1; i++ {
		p2 = p2.Next
	}
	if p2.Next == nil {
		p1 = p1.Next
		return p1
	}
	var pre *ListNode
	for p2.Next != nil {
		pre = p1
		p1 = p1.Next
		p2 = p2.Next
	}
	pre.Next = p1.Next
	return head
}
