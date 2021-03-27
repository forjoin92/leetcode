package leetcode_61

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_rotateRight1(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	k := 2
	expected := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
					},
				},
			},
		},
	}
	actual := rotateRight(head, k)
	assert.Equal(t, expected, actual)
}

func Test_rotateRight2(t *testing.T) {
	head := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 2,
			},
		},
	}
	k := 4
	expected := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 1,
			},
		},
	}
	actual := rotateRight(head, k)
	assert.Equal(t, expected, actual)
}
