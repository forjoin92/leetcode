package leetcode_69

func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	var res int
	left, right, y := 0, x, x
	for left <= right {
		middle := (left + right) / 2
		if middle == y/middle {
			return middle
		} else if middle > y/middle {
			right = middle - 1
		} else {
			left = middle + 1
			res = middle
		}
	}
	return res
}
