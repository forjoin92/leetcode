package leetcode_79

func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	var res bool
	for i := 0; i < len(board); i++ {
		visited[i] = make([]bool, len(board[0]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] != word[0] {
				continue
			}
			dfs(&res, board, i, j, word, 0, visited)
		}
	}
	return res
}

func dfs(res *bool, board [][]byte, x, y int, word string, index int, visited [][]bool) {
	if *res {
		return
	}
	if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) {
		return
	}
	if visited[x][y] {
		return
	}
	if index == len(word)-1 {
		if board[x][y] == word[index] {
			*res = true
		}
		return
	}
	if board[x][y] != word[index] {
		return
	}

	visited[x][y] = true
	dfs(res, board, x, y+1, word, index+1, visited)
	dfs(res, board, x, y-1, word, index+1, visited)
	dfs(res, board, x+1, y, word, index+1, visited)
	dfs(res, board, x-1, y, word, index+1, visited)
	visited[x][y] = false
}
