package q19

import "math"

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if divisor == 1 {
		return dividend
	}
	if divisor == -1 {
		if dividend > math.MinInt32 {
			return -dividend
		}
		return math.MaxInt32
	}
	result := 1
	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		result = -1
	}
	n1, n2 := abs(dividend), abs(divisor)
	return result * div(n1, n2)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func div(dividend, divisor int) int {
	if dividend < divisor {
		return 0
	}
	count := 1
	tmp := divisor
	for dividend >= tmp+tmp {
		tmp += tmp
		count += count
	}
	return count + div(dividend-tmp, divisor)
}
