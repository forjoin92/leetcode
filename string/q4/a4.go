package q4

func isAnagram(s string, t string) bool {
	m := make(map[int32]int)
	for _, c := range s {
		m[c]++
	}
	for _, c := range t {
		m[c]--
	}
	for _, n := range m {
		if n != 0 {
			return false
		}
	}
	return true
}
