package leetcode_26

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	var res int
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		} else {
			nums[res] = nums[i]
		}
		res++
	}
	return res
}
