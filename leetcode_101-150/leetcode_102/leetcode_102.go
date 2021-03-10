package leetcode_102

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func levelOrder(root *TreeNode) [][]int {
// 	var (
// 		res      [][]int
// 		resSlice []int
// 	)
// 	if root == nil {
// 		return res
// 	}
// 	queue1 := list.New()
// 	queue2 := list.New()
// 	queue1.PushBack(root)
// 	for {
// 		ele := queue1.Front()
// 		node := ele.Value.(*TreeNode)
// 		queue1.Remove(ele)
// 		resSlice = append(resSlice, node.Val)
// 		if node.Left != nil {
// 			queue2.PushBack(node.Left)
// 		}
// 		if node.Right != nil {
// 			queue2.PushBack(node.Right)
// 		}
// 		if queue1.Len() == 0 {
// 			res = append(res, resSlice)
// 			resSlice = nil
// 			if queue2.Len() == 0 {
// 				break
// 			}
// 			queue1, queue2 = queue2, queue1
// 		}
// 	}
// 	return res
// }

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) != 0 {
		length := len(queue)
		var subRes []int
		for i := 0; i < length; i++ {
			subRes = append(subRes, queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		res = append(res, subRes)
		queue = queue[length:]
	}
	return res
}
