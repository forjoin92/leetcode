package leetcode_1006

func clumsy(N int) int {
	var (
		results   []int
		operators = []byte{'*', '/', '+'}
		index     = 0
		pre       = N
	)
	for n := N - 1; n > 0; n-- {
		if index > 2 {
			index = 0
		}
		op := operators[index]
		index++
		switch op {
		case '*':
			pre *= n
		case '/':
			pre /= n
		case '+':
			results = append(results, pre)
			results = append(results, n)
			pre = n - 1
			n--
		}
	}
	if pre != 0 {
		results = append(results, pre)
	}
	var res = results[0]
	for i, num := range results[1:] {
		if i%2 == 0 {
			res += num
		} else {
			res -= num
		}
	}
	return res
}
