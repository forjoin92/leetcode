package leetcode_36

func isValidSudoku(board [][]byte) bool {
	isValid := true
	for i := 0; i < len(board); i++ {
		if isValid = isValidRow(board[i]); !isValid {
			return false
		}
		if isValid = isValidColumn(board, i); !isValid {
			return false
		}
	}
	for i := 0; i < len(board); i += 3 {
		for j := 0; j < len(board[i]); j += 3 {
			if isValid = isValidCorner(board, i, j, i+2, j+2); !isValid {
				return false
			}
		}
	}
	return true
}

func isValidRow(row []byte) bool {
	m := make(map[byte]struct{})
	for _, num := range row {
		if num != '.' {
			if _, exists := m[num]; exists {
				return false
			}
			m[num] = struct{}{}
		}
	}
	return true
}

func isValidColumn(board [][]byte, column int) bool {
	m := make(map[byte]struct{})
	for i := 0; i < len(board); i++ {
		if board[i][column] != '.' {
			if _, exists := m[board[i][column]]; exists {
				return false
			}
			m[board[i][column]] = struct{}{}
		}
	}
	return true
}

func isValidCorner(board [][]byte, x1, y1, x2, y2 int) bool {
	m := make(map[byte]struct{})
	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			if board[i][j] != '.' {
				if _, exists := m[board[i][j]]; exists {
					return false
				}
				m[board[i][j]] = struct{}{}
			}
		}
	}
	return true
}
