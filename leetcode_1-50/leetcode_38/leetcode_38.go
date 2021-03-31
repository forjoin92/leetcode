package leetcode_38

import (
	"strconv"
)

func countAndSay(n int) string {
	res := "1"
	for i := 1; i < n; i++ {
		var newRes string
		char, count := res[0], 1
		for i := 1; i < len(res); i++ {
			if res[i] == res[i-1] {
				count++
			} else {
				newRes = newRes + strconv.Itoa(count) + string(char)
				char = res[i]
				count = 1
			}
		}
		newRes = newRes + strconv.Itoa(count) + string(char)
		res = newRes
	}
	return res
}
