package q3

func rotate(nums []int, k int) {
	if k <= 0 || len(nums) <= 1 || len(nums) == k {
		return
	}

	for n, from, v, reach := 0, 0, nums[0], 0; n < len(nums); n++ {
		to := from + k
		if to >= len(nums) && to%len(nums) == reach {
			to %= len(nums)
			from = to + 1
			reach++
			nums[to] = v
			v = nums[from]
		} else {
			if to >= len(nums) {
				to %= len(nums)
			}
			from = to
			nums[to], v = v, nums[to]
		}
	}
}
