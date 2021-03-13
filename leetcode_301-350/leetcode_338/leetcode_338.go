package leetcode_338

func countBits(num int) []int {
	res := make([]int, num+1)
	for n := 1; n <= num; n++ {
		res[n] = res[n&(n-1)] + 1
	}
	return res
}
