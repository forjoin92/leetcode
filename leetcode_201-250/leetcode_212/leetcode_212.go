package leetcode_212

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

func findWords(board [][]byte, words []string) []string {
	if len(words) == 0 {
		return nil
	}
	result := make(map[string]bool)
	trie := Constructor()
	for _, word := range words {
		trie.Insert(word)
	}
	visited := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		visited[i] = make([]bool, len(board[0]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			dfs(result, board, "", i, j, visited, trie)
		}
	}
	var res []string
	for word := range result {
		res = append(res, word)
	}
	return res
}

func dfs(result map[string]bool, board [][]byte, str string, x, y int, visited [][]bool, trie Trie) {
	if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) {
		return
	}
	if visited[x][y] {
		return
	}
	str += string(board[x][y])
	if !trie.StartsWith(str) {
		return
	}
	if trie.Search(str) {
		result[str] = true
	}
	visited[x][y] = true
	dfs(result, board, str, x, y+1, visited, trie)
	dfs(result, board, str, x, y-1, visited, trie)
	dfs(result, board, str, x+1, y, visited, trie)
	dfs(result, board, str, x-1, y, visited, trie)
	visited[x][y] = false
}
