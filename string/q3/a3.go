package q3

func firstUniqChar(s string) int {
	m := make(map[int32]int)
	for _, c := range s {
		m[c]++
	}
	for i, c := range s {
		if m[c] == 1 {
			return i
		}
	}
	return -1
}
