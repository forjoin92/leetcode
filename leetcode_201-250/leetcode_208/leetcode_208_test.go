package leetcode_208

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Trie(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple")
	assert.Equal(t, true, trie.Search("apple"))
	assert.Equal(t, false, trie.Search("app"))
	assert.Equal(t, true, trie.StartsWith("app"))
	trie.Insert("app")
	assert.Equal(t, true, trie.Search("app"))
}
