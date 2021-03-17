package leetcode_70

func climbStairs(n int) int {
	if n < 2 {
		return 1
	}
	pre1, pre2, res := 1, 1, 0
	for i := 2; i <= n; i++ {
		res = pre1 + pre2
		pre1 = pre2
		pre2 = res
	}
	return res
}
