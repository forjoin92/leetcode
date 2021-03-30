package leetcode_25

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}
	var (
		res, pre      *ListNode
		cur, start, n = head, head, 0
	)
	for cur != nil {
		n++
		if n == k {
			if res == nil {
				res = cur
			}
			temp := cur.Next
			reverse(start, k)
			if pre != nil {
				pre.Next = cur
			}
			pre = start
			pre.Next, start, cur, n = temp, temp, temp, 0
		} else {
			cur = cur.Next
		}
	}
	return res
}

func reverse(head *ListNode, k int) {
	var (
		cur = head
		pre *ListNode
	)
	for i := 0; i < k; i++ {
		cur, cur.Next, pre = cur.Next, pre, cur
	}
}
