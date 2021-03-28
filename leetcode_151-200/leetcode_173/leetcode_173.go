package leetcode_173

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	trees []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	var bSTIterator BSTIterator
	InOrder(root, &bSTIterator.trees)
	return bSTIterator
}

func (this *BSTIterator) Next() int {
	res := this.trees[0]
	this.trees = this.trees[1:]
	return res.Val
}

func (this *BSTIterator) HasNext() bool {
	if len(this.trees) != 0 {
		return true
	}
	return false
}

func InOrder(root *TreeNode, trees *[]*TreeNode) {
	if root.Left != nil {
		InOrder(root.Left, trees)
	}
	*trees = append(*trees, root)
	if root.Right != nil {
		InOrder(root.Right, trees)
	}
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
