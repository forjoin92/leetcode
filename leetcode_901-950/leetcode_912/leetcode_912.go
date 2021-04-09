package leetcode_912

// func sortArray(nums []int) []int {
// 	if len(nums) < 2 {
// 		return nums
// 	}
// 	mid := len(nums) / 2
// 	return merge(sortArray(nums[:mid]), sortArray(nums[mid:]))
// }

// func merge(nums1, nums2 []int) []int {
// 	a1, a2 := 0, 0
// 	var res []int
// 	for a1 < len(nums1) || a2 < len(nums2) {
// 		var temp int
// 		if a1 == len(nums1) {
// 			temp = nums2[a2]
// 			a2++
// 		} else if a2 == len(nums2) {
// 			temp = nums1[a1]
// 			a1++
// 		} else {
// 			if nums1[a1] < nums2[a2] {
// 				temp = nums1[a1]
// 				a1++
// 			} else {
// 				temp = nums2[a2]
// 				a2++
// 			}
// 		}
// 		res = append(res, temp)
// 	}
// 	return res
// }

func sortArray(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	return quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, left, right int) []int {
	l, r := left, right
	if l < r {
		index := partition(nums, l, r)
		quickSort(nums, l, index-1)
		quickSort(nums, index+1, r)
	}
	return nums
}

func partition(nums []int, left, right int) int {
	pivot, index := left, left+1
	for i := index; i <= right; i++ {
		if nums[i] < nums[pivot] {
			nums[index], nums[i] = nums[i], nums[index]
			index++
		}
	}
	nums[pivot], nums[index-1] = nums[index-1], nums[pivot]
	return index - 1
}
