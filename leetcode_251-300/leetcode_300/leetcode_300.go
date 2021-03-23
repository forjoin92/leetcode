package leetcode_300

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxArr := make([]int, len(nums))
	res := 1
	for i := 0; i < len(nums); i++ {
		maxArr[i] = 1
	}
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				maxArr[i] = max(maxArr[i], maxArr[j]+1)
			}
		}
		res = max(res, maxArr[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
