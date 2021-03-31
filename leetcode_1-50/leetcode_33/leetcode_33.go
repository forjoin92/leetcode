package leetcode_33

func search(nums []int, target int) int {
	if len(nums) == 1 && nums[0] != target {
		return -1
	}
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			if nums[0] > nums[len(nums)-1] {
				if nums[left] <= nums[mid] {
					left = mid + 1
				} else {
					if nums[left] == target {
						return left
					} else if nums[left] < target {
						right = mid - 1
					} else {
						left = mid + 1
					}
				}
			} else {
				left = mid + 1
			}
		} else {
			if nums[0] > nums[len(nums)-1] {
				if nums[left] <= nums[mid] {
					if nums[left] == target {
						return left
					} else if nums[left] < target {
						right = mid - 1
					} else {
						left = mid + 1
					}
				} else {
					right = mid - 1
				}
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
