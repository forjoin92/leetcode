package leetcode_24

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	s := new(ListNode)
	pre := s
	pre.Next = head
	for pre.Next != nil && pre.Next.Next != nil {
		a := pre.Next
		b := a.Next
		a.Next, b.Next, pre.Next = b.Next, a, b
		pre = a
	}
	return s.Next
}
