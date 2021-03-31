package leetcode_977

func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	left, right, pos := 0, len(nums)-1, len(nums)-1
	for pos > -1 {
		if nums[left]*nums[left] > nums[right]*nums[right] {
			res[pos] = nums[left] * nums[left]
			left++
		} else {
			res[pos] = nums[right] * nums[right]
			right--
		}
		pos--
	}
	return res
}
