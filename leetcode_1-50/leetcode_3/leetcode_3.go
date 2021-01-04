package leetcode_3

func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	mp := make(map[int32]int)
	start, end := 0, 1
	res := 0
	for i, c := range s {
		if i > end {
			end = i
		}
		if v, exists := mp[c]; exists {
			if res < end-start {
				res = end - start
			}
			if v >= start {
				start = v + 1
			}
		}
		mp[c] = i
	}
	if res < end-start+1 {
		res = end - start + 1
	}
	return res
}
