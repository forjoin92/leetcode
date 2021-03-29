package leetcode_23

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	var heads []*ListNode
	for _, head := range lists {
		if head != nil {
			heads = append(heads, head)
		}
	}
	var res, cur *ListNode
	for len(heads) != 0 {
		index := min(heads)
		if cur == nil {
			cur = heads[index]
			res = cur
		} else {
			cur.Next = heads[index]
			cur = cur.Next
		}
		if heads[index].Next != nil {
			heads[index] = heads[index].Next
		} else {
			heads = append(heads[:index], heads[index+1:]...)
		}
	}
	return res
}

func min(lists []*ListNode) int {
	res := 0
	for i, node := range lists[1:] {
		if node == nil {
			continue
		}
		if node.Val < lists[res].Val {
			res = i + 1
		}
	}
	return res
}
