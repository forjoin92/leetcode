package q16

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	mp := map[int32][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}

	var res []string
	for _, num := range digits {
		str := mp[num]
		var newRes []string
		if len(res) == 0 {
			res = str
			continue
		}
		for _, r := range res {
			for _, s := range str {
				newRes = append(newRes, r+s)
			}
		}
		res = newRes
	}
	return res
}
