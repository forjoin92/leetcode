package leetcode_8

import (
	"math"
	"unicode"
)

func myAtoi(str string) int {
	str1 := trimLeft(str)
	if len(str1) == 0 {
		return 0
	}
	var res int64
	negative := int64(1)
	if str1[0] == '-' {
		negative = -1
		str1 = str1[1:]
	} else if str1[0] == '+' {
		str1 = str1[1:]
	}
	for i := 0; i < len(str1); i++ {
		if !unicode.IsDigit(rune(str1[i])) {
			break
		}
		res = res*10 + int64(str1[i]-'0')
		if res*negative > math.MaxInt32 {
			return math.MaxInt32
		}
		if res*negative < math.MinInt32 {
			return math.MinInt32
		}
	}
	return int(res * negative)
}

func trimLeft(str string) string {
	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			return str[i:]
		}
	}
	return ""
}
