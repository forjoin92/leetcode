package leetcode_237

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_deleteNode(t *testing.T) {
	head := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 9,
				},
			},
		},
	}
	node := head.Next
	expected := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 9,
			},
		},
	}
	deleteNode(node)
	assert.Equal(t, expected, head)
}

func Test_deleteNode1(t *testing.T) {
	head := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 9,
				},
			},
		},
	}
	node := head.Next.Next
	expected := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 9,
			},
		},
	}
	deleteNode(node)
	assert.Equal(t, expected, head)
}
