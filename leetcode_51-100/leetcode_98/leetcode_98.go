package leetcode_98

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return check(root, nil, nil)
}

func check(root *TreeNode, max, min *TreeNode) bool {
	if root == nil {
		return true
	}
	if max != nil && root.Val >= max.Val {
		return false
	}
	if min != nil && root.Val <= min.Val {
		return false
	}
	return check(root.Left, root, min) && check(root.Right, max, root)
}
