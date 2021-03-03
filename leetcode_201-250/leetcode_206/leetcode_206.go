package leetcode_206

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var (
		cur = head
		pre *ListNode
	)
	for cur != nil {
		cur, cur.Next, pre = cur.Next, pre, cur
	}
	return pre
}
