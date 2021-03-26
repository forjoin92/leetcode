package leetcode_83

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_deleteDuplicates1(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
			},
		},
	}
	expected := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}
	actual := deleteDuplicates(head)
	assert.Equal(t, expected, actual)
}

func Test_deleteDuplicates2(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
		},
	}
	expected := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	actual := deleteDuplicates(head)
	assert.Equal(t, expected, actual)
}
