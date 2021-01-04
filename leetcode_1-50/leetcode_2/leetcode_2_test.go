package leetcode_2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_addTwoNumbers(t *testing.T) {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	expected := &ListNode{
		Val: 7,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 8,
			},
		},
	}
	res := addTwoNumbers(l1, l2)
	assert.Equal(t, expected, res)
}

func Test_addTwoNumbers1(t *testing.T) {
	l1 := &ListNode{
		Val: 5,
	}
	l2 := &ListNode{
		Val: 5,
	}
	expected := &ListNode{
		Val: 0,
		Next: &ListNode{
			Val: 1,
		},
	}
	res := addTwoNumbers(l1, l2)
	assert.Equal(t, expected, res)
}
