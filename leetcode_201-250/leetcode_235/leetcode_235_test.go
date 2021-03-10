package leetcode_235

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lowestCommonAncestor1(t *testing.T) {
	head := &TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 5,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 7,
			},
			Right: &TreeNode{
				Val: 9,
			},
		},
	}
	expected := head
	res := lowestCommonAncestor(head, head.Left, head.Right)
	assert.Equal(t, expected, res)
}

func Test_lowestCommonAncestor2(t *testing.T) {
	head := &TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 5,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 7,
			},
			Right: &TreeNode{
				Val: 9,
			},
		},
	}
	expected := head.Left
	res := lowestCommonAncestor(head, head.Left, head.Left.Right)
	assert.Equal(t, expected, res)
}
