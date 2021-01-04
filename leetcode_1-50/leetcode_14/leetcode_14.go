package leetcode_14

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	str0 := strs[0]
	res := ""
	for i := 0; i < len(str0); i++ {
		for j := 1; j < len(strs); j++ {
			if len(strs[j]) <= i {
				return res
			}
			if str0[i] != strs[j][i] {
				return res
			}
		}
		res += string(str0[i])
	}
	return res
}
