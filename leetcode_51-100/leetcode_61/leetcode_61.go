package leetcode_61

type ListNode struct {
	Val  int
	Next *ListNode
}

// func rotateRight(head *ListNode, k int) *ListNode {
// 	if head == nil || k == 0 {
// 		return head
// 	}
// 	p1, p2 := head, head
// 	for i := 0; i < k-1; i++ {
// 		if p2.Next == nil {
// 			p2 = head
// 		} else {
// 			p2 = p2.Next
// 		}
// 	}
// 	if p2.Next == nil {
// 		return head
// 	}
// 	var pre *ListNode
// 	for p2.Next != nil {
// 		pre = p1
// 		p1 = p1.Next
// 		p2 = p2.Next
// 	}
// 	p2.Next, pre.Next = head, nil
// 	return p1
// }

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}
	var dp []*ListNode
	p1, p2, l := head, head, 0
	for i := 0; i < k-1; i++ {
		l++
		dp = append(dp, p2)
		if p2.Next == nil {
			p2 = dp[(k-1)%l]
			break
			// p2 = head
		} else {
			p2 = p2.Next
		}
	}
	if p2.Next == nil {
		return head
	}
	var pre *ListNode
	for p2.Next != nil {
		pre = p1
		p1 = p1.Next
		p2 = p2.Next
	}
	p2.Next, pre.Next = head, nil
	return p1
}
