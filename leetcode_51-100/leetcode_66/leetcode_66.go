package leetcode_66

func plusOne(digits []int) []int {
	length := len(digits)
	carry := true
	for i := length - 1; i >= 0; i-- {
		if carry {
			if digits[i]+1 == 10 {
				carry = true
				digits[i] = 0
			} else {
				digits[i]++
				carry = false
			}
		}
	}
	if carry {
		return append([]int{1}, digits...)
	}
	return digits
}
