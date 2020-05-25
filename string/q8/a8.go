package q8

import "fmt"

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	res := "1"
	for i := 2; i <= n; i++ {
		res = parseStr(res)
	}
	return res
}

func parseStr(str string) string {
	var (
		res   string
		count int
		char  uint8
	)
	for i := 0; i < len(str); i++ {
		if count == 0 {
			char = str[i]
			count++
		} else {
			if char == str[i] {
				count++
			} else {
				res = fmt.Sprintf("%s%d%c", res, count, char)
				count = 1
				char = str[i]
			}
		}
	}
	if count > 0 {
		res = fmt.Sprintf("%s%d%c", res, count, char)
	}
	return res
}
