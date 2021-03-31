package leetcode_31

func nextPermutation(nums []int) {
	sortedPos := -1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			pos := find(nums[i], i+1, nums)
			nums[i], nums[pos] = nums[pos], nums[i]
			sortedPos = i + 1
			break
		}
	}
	if sortedPos == -1 {
		sortedPos = 0
	}
	for ; sortedPos < len(nums); sortedPos++ {
		min := sortedPos
		for i := sortedPos + 1; i < len(nums); i++ {
			if nums[i] < nums[min] {
				min = i
			}
		}
		if min != sortedPos {
			nums[min], nums[sortedPos] = nums[sortedPos], nums[min]
		}
	}
}

func find(n, pos int, nums []int) int {
	res := pos
	for i := pos + 1; i < len(nums); i++ {
		if nums[i] > n && nums[i] < nums[res] {
			res = i
		}
	}
	return res
}
