package q4

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var res, tail *ListNode
	if l1.Val < l2.Val {
		res = &ListNode{
			Val: l1.Val,
		}
		l1 = l1.Next
		tail = res
	} else if l2.Val < l1.Val {
		res = &ListNode{
			Val: l2.Val,
		}
		l2 = l2.Next
		tail = res
	} else {
		res = &ListNode{
			Val: l1.Val,
			Next: &ListNode{
				Val: l2.Val,
			},
		}
		l1 = l1.Next
		l2 = l2.Next
		tail = res.Next
	}
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tail.Next = &ListNode{
				Val: l1.Val,
			}
			l1 = l1.Next
			tail = tail.Next
		} else if l2.Val < l1.Val {
			tail.Next = &ListNode{
				Val: l2.Val,
			}
			l2 = l2.Next
			tail = tail.Next
		} else {
			tail.Next = &ListNode{
				Val: l1.Val,
				Next: &ListNode{
					Val: l2.Val,
				},
			}
			l1 = l1.Next
			l2 = l2.Next
			tail = tail.Next.Next
		}
	}
	if l1 != nil {
		tail.Next = l1
	}
	if l2 != nil {
		tail.Next = l2
	}
	return res
}
