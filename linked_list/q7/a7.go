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
	var l3, l4 *ListNode
	for l1 != nil || l2 != nil {
		num := 0
		if l1 == nil {
			num = l2.Val + carry
		} else if l2 == nil {
			num = l1.Val + carry
		} else {
			num = l1.Val + l2.Val + carry
		}
		if num >= 10 {
			carry = 1
			num = num - 10
		} else {
			carry = 0
		}
		if l4 == nil {
			l3 = &ListNode{
				Val:  num,
				Next: nil,
			}
			l4 = l3
		} else {
			l3.Next = &ListNode{
				Val:  num,
				Next: nil,
			}
			l3 = l3.Next
		}
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	if carry == 1 {
		l3.Next = &ListNode{
			Val:  1,
			Next: nil,
		}
	}
	return l4
}
