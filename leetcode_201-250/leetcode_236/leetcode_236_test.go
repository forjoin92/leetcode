package leetcode_236

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lowestCommonAncestor1(t *testing.T) {
	head := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 8,
			},
		},
	}
	expected := head
	res := lowestCommonAncestor(head, head.Left, head.Right)
	assert.Equal(t, expected, res)
}

func Test_lowestCommonAncestor2(t *testing.T) {
	head := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
		},
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 8,
			},
		},
	}
	expected := head.Left
	res := lowestCommonAncestor(head, head.Left, head.Left.Right.Right)
	assert.Equal(t, expected, res)
}
