package leetcode_12

func intToRoman(num int) string {
	mp := map[int]string{
		1: "I", 4: "IV", 5: "V", 9: "IX",
		10: "X", 40: "XL", 50: "L", 90: "XC",
		100: "C", 400: "CD", 500: "D", 900: "CM", 1000: "M",
	}
	var res string
	div := 1
	switch {
	case num >= 1000:
		div = 1000
	case num >= 100:
		div = 100
	case num >= 10:
		div = 10
	}
	for div != 0 {
		if num/div != 0 {
			n := num / div
			switch {
			case n < 4:
				for i := 0; i < n; i++ {
					res += mp[div]
				}
			case n == 4:
				res += mp[div*4]
			case n == 5:
				res += mp[div*5]
			case n < 9 && n > 5:
				res += mp[div*5]
				for i := 0; i < n-5; i++ {
					res += mp[div]
				}
			case n == 9:
				res += mp[div*9]
			}
		}
		num = num - (num/div)*div
		div /= 10
	}
	return res
}
