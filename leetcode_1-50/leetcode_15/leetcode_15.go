package leetcode_15

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		p := len(nums) - 1
		for j := i + 1; j < len(nums); j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for j < p && nums[j]+nums[p] > -nums[i] {
				p--
			}
			if j == p {
				break
			}
			if nums[i]+nums[j]+nums[p] == 0 {
				res = append(res, []int{nums[i], nums[j], nums[p]})
			}
		}
	}
	return res
}
