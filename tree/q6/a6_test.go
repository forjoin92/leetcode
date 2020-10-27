package q6

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_preorderTraversal(t *testing.T) {
	tree := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	expected := []int{1, 2, 3}
	res := preorderTraversal(tree)
	assert.Equal(t, expected, res)
}
