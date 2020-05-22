package q2

import (
	"math"
)

func reverse(x int) int {
	var res int64
	for x != 0 {
		res = res*10 + int64(x)%10
		x = x / 10
	}
	if res > math.MaxInt32 || res < math.MinInt32 {
		return 0
	}
	return int(res)
}
