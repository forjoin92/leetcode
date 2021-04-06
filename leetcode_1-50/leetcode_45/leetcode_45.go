package leetcode_45

// func jump(nums []int) int {
// 	if len(nums) < 2 {
// 		return 0
// 	}
// 	res := len(nums) - 1
// 	dfs(&res, nums, 0, 0)
// 	return res
// }

// func dfs(res *int, nums []int, index int, step int) {
// 	if *res == 1 {
// 		return
// 	}
// 	if index >= len(nums)-1 {
// 		if *res > step {
// 			*res = step
// 		}
// 		return
// 	}
// 	step++
// 	if nums[index] != 0 {
// 		dfs(res, nums, index+nums[index], step)
// 	}
// 	dfs(res, nums, index+1, step)
// }

func jump(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return 0
	}
	res := 0
	dfs(&res, nums, 0, nums[0])
	return res
}

func dfs(res *int, nums []int, from, to int) {
	*res++
	if to >= len(nums)-1 {
		return
	}
	nextPos := from + 1
	for i := from + 2; i <= to; i++ {
		if nextPos+nums[nextPos] < i+nums[i] {
			nextPos = i
		}
	}
	dfs(res, nums, nextPos, nextPos+nums[nextPos])
}
