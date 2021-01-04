package leetcode_1365

func smallerNumbersThanCurrent(nums []int) []int {
	newNums := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		num := 0
		for j := 0; j < len(nums); j++ {
			if i != j && nums[j] < nums[i] {
				num++
			}
		}
		newNums[i] = num
	}
	return newNums
}
