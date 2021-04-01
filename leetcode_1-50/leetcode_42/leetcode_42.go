package leetcode_42

func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	start, res, end, up := 0, 0, 1, false
	for i := 1; i < len(height); i++ {
		for j := i; j < len(height); j++ {
			if height[j] >= height[start] {
				end = j
				for k := start + 1; k < end; k++ {
					res = res + (height[start] - height[k])
				}
				start, i, up = end, end, false
				continue
			}
			if height[j] > height[j-1] {
				if up {
					if height[end] < height[j] {
						end = j
						up = true
					}
				} else {
					end = j
					up = true
				}
			}
		}
		if up {
			for k := start + 1; k < end; k++ {
				temp := min(height[start], height[end])
				a := temp - height[k]
				if a > 0 {
					res = res + a
				}
			}
			start, i, up = end, end, false
		} else {
			start = i
		}
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
