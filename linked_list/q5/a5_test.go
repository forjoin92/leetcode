package q5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isPalindrome(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}
	expected := false
	res := isPalindrome(head)
	assert.Equal(t, expected, res)
}

func Test_isPalindrome1(t *testing.T) {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
				},
			},
		},
	}
	expected := true
	res := isPalindrome(head)
	assert.Equal(t, expected, res)
}
