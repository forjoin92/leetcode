package leetcode_34

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	left, right, res := 0, len(nums)-1, []int{-1, -1}

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			l1, r1 := left, mid
			for l1 <= r1 {
				mid1 := l1 + (r1-l1)/2
				if nums[mid1] == target {
					if mid1 > 0 {
						if nums[mid1-1] != target {
							res[0] = mid1
							break
						} else {
							r1 = mid1 - 1
						}
					} else {
						res[0] = mid1
						break
					}
				} else if nums[l1] < target {
					l1 = mid1 + 1
				} else {
					r1 = mid1 - 1
				}
			}
			l2, r2 := mid, right
			for l2 <= r2 {
				mid2 := l2 + (r2-l2)/2
				if nums[mid2] == target {
					if mid2 < len(nums)-1 {
						if nums[mid2+1] != target {
							res[1] = mid2
							break
						} else {
							l2 = mid2 + 1
						}
					} else {
						res[1] = mid2
						break
					}
				} else if nums[l2] < target {
					l2 = mid2 + 1
				} else {
					r2 = mid2 - 1
				}
			}
			break
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return res
}
