package leetcode_82

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	s := new(ListNode)
	pre := s
	pre.Next = head
	for pre.Next != nil {
		a := pre.Next
		for a.Next != nil && a.Val == a.Next.Val {
			a = a.Next
		}
		if pre.Next == a {
			pre = a
		} else {
			pre.Next = a.Next
		}
	}
	return s.Next
}
