package leetcode_83

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	pre := head
	for pre != nil {
		a := pre
		for a.Next != nil && a.Val == a.Next.Val {
			a = a.Next
		}
		pre.Next = a.Next
		pre = a.Next
	}
	return head
}
