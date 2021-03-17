package leetcode_52

// func totalNQueens(n int) int {
// 	var res int
// 	cols := make(map[int]bool)
// 	pie := make(map[int]bool)
// 	na := make(map[int]bool)
// 	dfs(&res, cols, pie, na, n, 0, 0)
// 	return res
// }

// func dfs(res *int, cols, pie, na map[int]bool, n, row int, state int) {
// 	if row >= n {
// 		if state == n {
// 			*res++
// 		}
// 	}
// 	for i := 0; i < n; i++ {
// 		if cols[i] || pie[row+i] || na[row-i] {
// 			continue
// 		}
// 		cols[i], pie[row+i], na[row-i] = true, true, true
// 		dfs(res, cols, pie, na, n, row+1, state+1)
// 		delete(cols, i)
// 		delete(pie, row+i)
// 		delete(na, row-i)
// 	}
// }

func totalNQueens(n int) int {
	var res int
	dfs(&res, n, 0, 0, 0, 0)
	return res
}

func dfs(res *int, n, row, cols, pie, na int) {
	if row >= n {
		*res++
	}
	bits := (^(cols | pie | na)) & ((1 << n) - 1)
	for bits != 0 {
		p := bits & -bits
		dfs(res, n, row+1, cols|p, (pie|p)<<1, (na|p)>>1)
		bits = bits & (bits - 1)
	}
}
