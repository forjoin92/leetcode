package q7

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if haystack == "" {
		return -1
	}

	for i, c := range haystack {
		if c == int32(needle[0]) && i+len(needle) <= len(haystack) && haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}
