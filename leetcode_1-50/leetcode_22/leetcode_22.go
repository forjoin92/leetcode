package leetcode_22

// func generateParenthesis(n int) []string {
// 	if n == 0 {
// 		return []string{}
// 	}
// 	if n == 1 {
// 		return []string{"()"}
// 	}
// 	pre := generateParenthesis(n - 1)
// 	return generate(pre)
// }

// func generate(pre []string) []string {
// 	mp := make(map[string]struct{})
// 	var res []string
// 	for _, s := range pre {
// 		base := "(" + s
// 		for i := 1; i < len(base); i++ {
// 			newStr := base[:i] + ")" + base[i:]
// 			if _, ok := mp[newStr]; !ok {
// 				mp[newStr] = struct{}{}
// 				res = append(res, newStr)
// 			}
// 		}
// 	}
// 	return res
// }

func generateParenthesis(n int) []string {
	var list []string
	gen(&list, 0, 0, n, "")
	return list
}

func gen(list *[]string, left, right, n int, str string) {
	if left == n && right == n {
		*list = append(*list, str)
	}
	if left < n {
		gen(list, left+1, right, n, str+"(")
	}
	if left > right && right < n {
		gen(list, left, right+1, n, str+")")
	}
}
