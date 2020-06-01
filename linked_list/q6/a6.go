package q6

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	fast, slow := head.Next, head
	for fast != nil {
		if fast.Next == nil || fast.Next.Next == nil {
			return false
		} else {
			fast = fast.Next.Next
			slow = slow.Next
			if fast == slow {
				return true
			}
		}
	}
	return false
}
