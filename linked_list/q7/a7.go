package q7

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	carry := 0
	res := l1
	var pre *ListNode
	for l1 != nil && l2 != nil {
		value := l1.Val + l2.Val + carry
		if value >= 10 {
			carry = 1
			l1.Val = value - 10
		} else {
			carry = 0
			l1.Val = value
		}
		pre = l1
		l1 = l1.Next
		l2 = l2.Next
	}
	if l2 != nil {
		l1 = l2
	}
	for l1 != nil {
		pre.Next = l1
		value := l1.Val + carry
		if value == 10 {
			carry = 1
			l1.Val = 0
		} else {
			l1.Val = value
			return res
		}
		pre = l1
		l1 = l1.Next
	}
	if carry == 1 {
		pre.Next = &ListNode{
			Val: 1,
		}
	}
	return res
}
