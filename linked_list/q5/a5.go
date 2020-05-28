package q5

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	fast, slow := head.Next, head
	newHead := slow.Next
	head.Next = nil
	num := 0
	for fast.Next != nil {
		if fast.Next.Next == nil {
			num = 1
			fast = fast.Next
		} else {
			fast = fast.Next.Next
			tmpSlow := slow
			tmp := newHead.Next
			slow = newHead
			newHead = tmp
			slow.Next = tmpSlow
		}
	}
	if num == 1 {
		newHead = newHead.Next
	}
	for slow != nil && newHead != nil {
		if slow.Val != newHead.Val {
			return false
		}
		slow = slow.Next
		newHead = newHead.Next
	}
	if slow != nil || newHead != nil {
		return false
	}
	return true
}
