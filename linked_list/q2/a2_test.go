package q2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeNthFromEnd(t *testing.T) {
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
	n := 2
	expected := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 5,
				},
			},
		},
	}
	res := removeNthFromEnd(head, n)
	assert.Equal(t, expected, res)
}

func Test_removeNthFromEnd1(t *testing.T) {
	head := &ListNode{
		Val: 1,
	}
	n := 1
	var expected *ListNode
	res := removeNthFromEnd(head, n)
	assert.Equal(t, expected, res)
}
