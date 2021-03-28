package leetcode_173

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_majorityElement1(t *testing.T) {
	root := &TreeNode{
		Val: 7,
		Left: &TreeNode{
			Val: 3,
		},
		Right: &TreeNode{
			Val: 15,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 20,
			},
		},
	}
	bSTIterator := Constructor(root)
	assert.Equal(t, bSTIterator.Next(), 3)
	assert.Equal(t, bSTIterator.Next(), 7)
	assert.Equal(t, bSTIterator.HasNext(), true)
	assert.Equal(t, bSTIterator.Next(), 9)
	assert.Equal(t, bSTIterator.HasNext(), true)
	assert.Equal(t, bSTIterator.Next(), 15)
	assert.Equal(t, bSTIterator.HasNext(), true)
	assert.Equal(t, bSTIterator.Next(), 20)
	assert.Equal(t, bSTIterator.HasNext(), false)
}
