package q5

func singleNumber(nums []int) int {
	res := nums[0]
	for _, num := range nums[1:] {
		res ^= num
	}
	return res
}
