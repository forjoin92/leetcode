package q1

func removeDuplicates(nums []int) int {
	length := len(nums)
	if length < 2 {
		return length
	}
	position := 0
	for i := 1; i < length; i++ {
		if nums[i] != nums[position] {
			position++
			nums[position] = nums[i]
		}
	}
	return position + 1
}
