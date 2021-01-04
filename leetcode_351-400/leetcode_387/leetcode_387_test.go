package leetcode_387

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_firstUniqChar(t *testing.T) {
	s1 := "leetcode"
	expected1 := 0
	res1 := firstUniqChar(s1)
	assert.Equal(t, expected1, res1)

	s2 := "loveleetcode"
	expected2 := 2
	res2 := firstUniqChar(s2)
	assert.Equal(t, expected2, res2)
}
