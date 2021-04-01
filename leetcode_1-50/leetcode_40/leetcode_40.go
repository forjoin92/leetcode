package leetcode_40

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var res [][]int
	mp := make(map[int]struct{})
	for i := 0; i < len(candidates); i++ {
		// if i > 0 && candidates[i] == candidates[i-1] {
		// 	continue
		// }
		dfs(candidates, []int{}, i, target-candidates[i], &res, mp)
	}
	return res
}

func dfs(candidates []int, result []int, index int, target int, res *[][]int, mp map[int]struct{}) {
	if target < 0 {
		return
	}
	result = append(result, candidates[index])
	if target == 0 {
		key := 0
		for _, v := range result {
			key = key*10 + v
		}
		if _, ok := mp[key]; ok {
			return
		}
		mp[key] = struct{}{}
		tmp := make([]int, len(result))
		copy(tmp, result)
		*res = append(*res, tmp)

		return
	}
	for i := 0; i < len(candidates); i++ {
		if i > index {
			dfs(candidates, result, i, target-candidates[i], res, mp)
		}
	}
}
