package leetcode_111

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minDepth1(t *testing.T) {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	expected := 2
	res := minDepth(root)
	assert.Equal(t, expected, res)
}

func Test_minDepth2(t *testing.T) {
	root := &TreeNode{
		Val: 2,
		Right: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 4,
				Right: &TreeNode{
					Val: 5,
					Right: &TreeNode{
						Val: 6,
					},
				},
			},
		},
	}
	expected := 5
	res := minDepth(root)
	assert.Equal(t, expected, res)
}
