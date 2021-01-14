package leetcode_1018

func prefixesDivBy5(A []int) []bool {
	res := make([]bool, 0, len(A))
	var (
		sum int
		f   bool
	)

	for i := 0; i < len(A); i++ {
		sum, f = check(A[:i+1], i+1, sum)
		res = append(res, f)
	}
	return res
}

func check(A []int, length int, pre int) (int, bool) {
	var sum int
	sum = pre*2 + A[length-1]
	if sum%5 == 0 {
		return 0, true
	}
	return sum % 10, false
}
