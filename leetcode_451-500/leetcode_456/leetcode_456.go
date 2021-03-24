package leetcode_456

// 循环，O(n*n*n)
// func find132pattern(nums []int) bool {
// 	if len(nums) < 3 {
// 		return false
// 	}
// 	for i := 0; i < len(nums); i++ {
// 		for k := i + 2; k < len(nums); k++ {
// 			if nums[i] >= nums[k] {
// 				continue
// 			}
// 			for j := i + 1; j < k; j++ {
// 				if nums[j] > nums[k] {
// 					return true
// 				} else {
// 					continue
// 				}
// 			}
// 		}
// 	}
// 	return false
// }

// dfs, O(n*n*n)
// func find132pattern(nums []int) bool {
// 	if len(nums) < 3 {
// 		return false
// 	}
// 	return dfs(0, 1, 2, nums)
// }

// func dfs(i, j, k int, nums []int) bool {
// 	if i >= j || j >= k || i >= len(nums)-1 || j >= len(nums)-1 || k > len(nums)-1 {
// 		return false
// 	}
// 	if nums[j] > nums[k] && nums[k] > nums[i] {
// 		return true
// 	}
// 	return dfs(i+1, j, k, nums) || dfs(i, j+1, k, nums) || dfs(i, j, k+1, nums)
// }

// dp，O(n*n)
func find132pattern(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	dp := make([]int, len(nums))
	min := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[min] {
			dp[i] = min
		} else {
			min = i
			dp[i] = -1
		}
	}
	for j, v := range dp[1:] {
		if v >= 0 {
			for k := j + 1 + 1; k < len(nums); k++ {
				if nums[k] > nums[v] && nums[k] < nums[j+1] {
					return true
				}
			}
		}
	}
	return false
}

// func find132pattern(nums []int) bool {
// 	if len(nums) < 3 {
// 		return false
// 	}
// 	dpMin := make([]int, len(nums))
// 	dpMax := make([]int, len(nums))
// 	min, max := 0, nums[0]
// 	dpMin[0], dpMax[0] = min, max
// 	for i := 1; i < len(nums); i++ {
// 		if nums[i] > nums[min] {
// 			dpMin[i] = min
// 			if nums[i] < max {
// 				dpMax[i] = max
// 			} else {
// 				max = nums[i]
// 				dpMax[i] = max
// 			}
// 		} else {
// 			min = i
// 			dpMin[i] = -1
// 			max = nums[min]
// 			dpMax[i] = max
// 		}

// 	}
// 	for i := 2; i < len(nums); i++ {
// 		if dpMin[i-1] == -1 {
// 			continue
// 		}
// 		if nums[i] < dpMax[i-1] && nums[i] > nums[dpMin[i-1]] {
// 			return true
// 		}
// 	}
// 	fmt.Println(dpMin)
// 	fmt.Println(dpMax)
// 	return false
// }
