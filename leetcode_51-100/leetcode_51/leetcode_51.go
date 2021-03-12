package leetcode_51

func solveNQueens(n int) [][]string {
	var result [][]string
	cols := make(map[int]bool)
	pie := make(map[int]bool)
	na := make(map[int]bool)
	dfs(&result, cols, pie, na, n, 0, []int{})
	return result
}

func dfs(result *[][]string, cols, pie, na map[int]bool, n, row int, state []int) {
	if row >= n {
		var item []string
		for _, col := range state {
			str := ""
			for i := 0; i < col; i++ {
				str += "."
			}
			str += "Q"
			for i := 0; i < n-col-1; i++ {
				str += "."
			}
			item = append(item, str)
		}
		*result = append(*result, item)
	}
	for i := 0; i < n; i++ {
		if cols[i] || pie[row+i] || na[row-i] {
			continue
		}
		cols[i], pie[row+i], na[row-i] = true, true, true
		dfs(result, cols, pie, na, n, row+1, append(state, i))
		delete(cols, i)
		delete(pie, row+i)
		delete(na, row-i)
	}
}
