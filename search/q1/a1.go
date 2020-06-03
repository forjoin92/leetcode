package q1

func firstBadVersion(n int) int {
	start, end := 1, n
	for start <= end {
		pos := (start + end) / 2
		if isBadVersion(pos) {
			if pos != 1 && isBadVersion(pos-1) {
				end = pos
			} else {
				return pos
			}
		} else {
			if start == pos {
				start++
			} else {
				start = pos
			}
		}
	}
	return 1
}

func isBadVersion(version int) bool {
	return false
}
