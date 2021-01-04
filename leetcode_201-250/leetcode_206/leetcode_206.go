package leetcode_206

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	pre, next := head, head.Next
	head.Next = nil
	for next != nil {
		n1 := next.Next
		next.Next = pre
		pre = next
		next = n1
	}
	return pre
}
