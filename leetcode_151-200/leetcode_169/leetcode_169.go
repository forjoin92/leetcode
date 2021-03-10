package leetcode_169

func majorityElement(nums []int) int {
	mp := make(map[int]int)
	for _, num := range nums {
		mp[num]++
	}
	for k, v := range mp {
		if v > len(nums)/2 {
			return k
		}
	}
	return 0
}
