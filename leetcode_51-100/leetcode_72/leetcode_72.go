package leetcode_72

func minDistance(word1 string, word2 string) int {
	len1, len2 := len(word1)+1, len(word2)+1
	dp := make([][]int, len1)
	for i := 0; i < len1; i++ {
		dp[i] = make([]int, len2)
	}
	for i := 0; i < len1; i++ {
		dp[i][0] = i
	}
	for j := 0; j < len2; j++ {
		dp[0][j] = j
	}
	for i := 1; i < len1; i++ {
		for j := 1; j < len2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])
			}
		}
	}
	return dp[len1-1][len2-1]
}

func min(a, b, c int) int {
	res := a
	if res > b {
		res = b
	}
	if res > c {
		res = c
	}
	return res
}