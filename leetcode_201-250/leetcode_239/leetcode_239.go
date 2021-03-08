package leetcode_239

func maxSlidingWindow(nums []int, k int) []int {
	var window, res []int
	for index, num := range nums {
		if index >= k && window[0] <= index-k {
			window = window[1:]
		}
		for len(window) != 0 && nums[window[len(window)-1]] <= num {
			window = window[:len(window)-1]
		}
		window = append(window, index)
		if index >= k-1 {
			res = append(res, nums[window[0]])
		}
	}
	return res
}
