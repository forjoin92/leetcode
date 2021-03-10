package leetcode_15

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		head, tail := i+1, len(nums)-1
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for head < tail {
			n1, n2 := nums[head], nums[tail]
			total := nums[i] + n1 + n2
			if total > 0 {
				tail--
			} else if total < 0 {
				head++
			} else {
				res = append(res, []int{nums[i], nums[head], nums[tail]})
				for head < tail && nums[head] == n1 {
					head++
				}
				for head < tail && nums[tail] == n2 {
					tail--
				}
			}
		}
	}
	return res
}
