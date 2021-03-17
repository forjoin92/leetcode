package leetcode_120

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	minArr := triangle[len(triangle)-1]
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			minArr[j] = triangle[i][j] + min(minArr[j], minArr[j+1])
		}
	}
	return minArr[0]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
