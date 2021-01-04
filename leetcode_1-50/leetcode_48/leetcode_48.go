package leetcode_48

func rotate(matrix [][]int) {
	l := len(matrix)
	for x := 0; x < l/2; x++ {
		for y := x; y < l-1-x; y++ {
			temp := matrix[x][y]
			matrix[x][y] = matrix[l-1-y][x]
			matrix[l-1-y][x] = matrix[l-1-x][l-1-y]
			matrix[l-1-x][l-1-y] = matrix[l-1-(l-1-y)][l-1-x]
			matrix[l-1-(l-1-y)][l-1-x] = temp
		}
	}
}
