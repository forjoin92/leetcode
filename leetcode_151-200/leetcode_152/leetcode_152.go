package leetcode_152

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var (
		max [2][2]int
		res int
	)
	max[0][0], max[0][1], res = nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		x, y := i%2, (i-1)%2
		max[x][0] = maxOf3(max[y][0]*nums[i], max[y][1]*nums[i], nums[i])
		max[x][1] = minOf3(max[y][0]*nums[i], max[y][1]*nums[i], nums[i])
		if res < max[x][0] {
			res = max[x][0]
		}
	}
	return res
}

func maxOf3(a, b, c int) int {
	res := a
	if res < b {
		res = b
	}
	if res < c {
		res = c
	}
	return res
}

func minOf3(a, b, c int) int {
	res := a
	if res > b {
		res = b
	}
	if res > c {
		res = c
	}
	return res
}
