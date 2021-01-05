package leetcode_830

func largeGroupPositions(s string) [][]int {
	var start, end int
	var res [][]int
	for i := 1; i < len(s); i++ {
		if s[start] != s[i] {
			if end-start >= 2 {
				res = append(res, []int{start, end})
			}
			start = i
		}
		end = i
	}
	if end-start >= 2 {
		res = append(res, []int{start, end})
	}
	return res
}
