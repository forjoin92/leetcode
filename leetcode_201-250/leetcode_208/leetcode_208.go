package leetcode_208

type Trie struct {
	end      bool
	chridren [26]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	for _, c := range word {
		if this.chridren[c-'a'] == nil {
			this.chridren[c-'a'] = &Trie{}
		}
		this = this.chridren[c-'a']
	}
	this.end = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	for _, c := range word {
		if this.chridren[c-'a'] == nil {
			return false
		}
		this = this.chridren[c-'a']
	}
	if this.end {
		return true
	}
	return false
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for _, c := range prefix {
		if this.chridren[c-'a'] == nil {
			return false
		}
		this = this.chridren[c-'a']
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
