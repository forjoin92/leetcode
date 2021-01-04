package leetcode_22

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	if n == 1 {
		return []string{"()"}
	}
	pre := generateParenthesis(n - 1)
	return generate(pre)
}

func generate(pre []string) []string {
	mp := make(map[string]struct{})
	var res []string
	for _, s := range pre {
		base := "(" + s
		for i := 1; i < len(base); i++ {
			newStr := base[:i] + ")" + base[i:]
			if _, ok := mp[newStr]; !ok {
				mp[newStr] = struct{}{}
				res = append(res, newStr)
			}
		}
	}
	return res
}
