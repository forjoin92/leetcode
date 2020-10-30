package q14

func romanToInt(s string) int {
	mp := make(map[rune]int)
	mp['I'] = 1
	mp['V'] = 5
	mp['X'] = 10
	mp['L'] = 50
	mp['C'] = 100
	mp['D'] = 500
	mp['M'] = 1000
	res := 0
	pre := 0
	for _, c := range s {
		if mp[c] > pre {
			res = res - pre + mp[c] - pre
			pre = mp[c]
		} else {
			res += mp[c]
			pre = mp[c]
		}
	}
	return res
}
