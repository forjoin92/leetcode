package leetcode_28

func strStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}
	for i := 0; i < len(haystack)-(len(needle)-1); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}
