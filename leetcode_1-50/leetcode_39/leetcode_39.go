package leetcode_39

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	for i := 0; i < len(candidates); i++ {
		dfs(candidates, []int{}, i, target-candidates[i], &res)
	}
	return res
}

func dfs(candidates []int, result []int, index int, target int, res *[][]int) {
	if target < 0 {
		return
	}
	result = append(result, candidates[index])
	if target == 0 {
		tmp := make([]int, len(result))
		copy(tmp, result)
		*res = append(*res, tmp)
		return
	}
	for i := 0; i < len(candidates); i++ {
		if i >= index {
			dfs(candidates, result, i, target-candidates[i], res)
		}
	}
}
