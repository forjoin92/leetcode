package leetcode_80

func removeDuplicates(nums []int) int {
	c1, cur, count := -1, nums[0], 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == cur {
			if count == 2 && c1 == -1 {
				c1 = i
			}
			count++
		} else {
			if nums[i] != cur {
				cur = nums[i]
				count = 1
			}
		}
		if c1 != -1 && count < 3 && nums[c1] != nums[i] {
			nums[c1], nums[i] = nums[i], nums[c1]
			c1++
		}
	}
	if c1 == -1 {
		return len(nums)
	}
	return c1
}
