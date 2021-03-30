package leetcode_25

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseKGroup1(t *testing.T) {
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
		Val: 2,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	actual := reverseKGroup(head, k)
	assert.Equal(t, expected, actual)
}

func Test_reverseKGroup2(t *testing.T) {
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
	k := 3
	expected := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	actual := reverseKGroup(head, k)
	assert.Equal(t, expected, actual)
}

func Test_reverseKGroup3(t *testing.T) {
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
	k := 1
	expected := &ListNode{
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
	actual := reverseKGroup(head, k)
	assert.Equal(t, expected, actual)
}

func Test_reverseKGroup4(t *testing.T) {
	head := &ListNode{
		Val: 1,
	}
	k := 1
	expected := &ListNode{
		Val: 1,
	}
	actual := reverseKGroup(head, k)
	assert.Equal(t, expected, actual)
}
