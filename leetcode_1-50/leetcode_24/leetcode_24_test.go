package leetcode_24

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_swapPairs1(t *testing.T) {
	l := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
	}
	expected := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 3,
				},
			},
		},
	}
	actual := swapPairs(l)
	assert.Equal(t, expected, actual)
}

func Test_swapPairs2(t *testing.T) {
	l := new(ListNode)
	expected := new(ListNode)
	actual := swapPairs(l)
	assert.Equal(t, expected, actual)
}

func Test_swapPairs3(t *testing.T) {
	l := &ListNode{
		Val: 1,
	}
	expected := &ListNode{
		Val: 1,
	}
	actual := swapPairs(l)
	assert.Equal(t, expected, actual)
}
