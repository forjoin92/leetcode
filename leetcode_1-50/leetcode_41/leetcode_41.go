package leetcode_41

import "sort"

func firstMissingPositive(nums []int) int {
	if len(nums) < 0 {
		return 1
	}
	if len(nums) == 1 {
		if nums[0] != 1 {
			return 1
		} else {
			return 2
		}
	}
	sort.Ints(nums)
	res := 1
	for i := 0; i < len(nums); i++ {
		if nums[i] <= 0 {
			continue
		}
		if i == 0 {
			if nums[i] != 1 {
				return 1
			} else {
				res++
			}
		} else {
			if nums[i] == nums[i-1] {
				continue
			}
			if nums[i] == res {
				res++
			}
		}
	}
	return res
}
