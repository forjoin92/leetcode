package q9

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	res := strs[0]
	var base int
	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < len(res) {
			base = len(strs[i])
		} else {
			base = len(res)
		}
		if base == 0 {
			return ""
		}
		res = res[:base]
		for j := 0; j < base; j++ {
			if res[j] != strs[i][j] {
				res = res[:j]
				break
			}
		}
	}
	return res
}
