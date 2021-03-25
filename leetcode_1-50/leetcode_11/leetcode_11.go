package leetcode_11

// O(n*n)，超时了
// func maxArea(height []int) int {
// 	res := 0
// 	for i := 0; i < len(height); i++ {
// 		if height[i] == 0 {
// 			continue
// 		}
// 		for j := i + 1; j < len(height); j++ {
// 			if height[j] == 0 {
// 				continue
// 			}
// 			if min(height[i], height[j])*(j-i) > res {
// 				res = min(height[i], height[j]) * (j - i)
// 			}
// 		}
// 	}
// 	return res
// }

// func min(a, b int) int {
// 	if a > b {
// 		return b
// 	}
// 	return a
// }

// O(n)
func maxArea(height []int) int {
	res, l, r := 0, 0, len(height)-1
	for l < r {
		res = max(res, min(height[l], height[r])*(r-l))
		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
