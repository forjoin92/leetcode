package q5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sortedArrayToBST(t *testing.T) {
	nums := []int{-10, -3, 0, 5, 9}
	expected := &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: -3,
			Left: &TreeNode{
				Val: -10,
			},
		},
		Right: &TreeNode{
			Val: 9,
			Left: &TreeNode{
				Val: 5,
			},
		},
	}
	res := sortedArrayToBST(nums)
	assert.Equal(t, expected, res)
}
