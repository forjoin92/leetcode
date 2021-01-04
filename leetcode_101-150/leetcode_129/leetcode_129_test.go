package leetcode_129

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sumNumbers1(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	expected := 25
	res := sumNumbers(tree)
	assert.Equal(t, expected, res)
}

func Test_sumNumbers2(t *testing.T) {
	tree := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val: 9,
			Left: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   0,
			Left:  nil,
			Right: nil,
		},
	}
	expected := 1026
	res := sumNumbers(tree)
	assert.Equal(t, expected, res)
}
