package leetcode_90

func subsetsWithDup(nums []int) [][]int {
	mp := make(map[int]int)
	var res [][]int
	for _, num := range nums {
		mp[num]++
	}
	for num, count := range mp {
		numArr := [][]int{
			{num},
		}
		for i := 1; i < count; i++ {
			tmp := append(numArr[len(numArr)-1], num)
			numArr = append(numArr, tmp)
		}
		var newRes [][]int
		for _, arr := range numArr {
			resLen := len(res)
			for j := 0; j < resLen; j++ {
				resJ := make([]int, len(res[j]))
				copy(resJ, res[j])
				temp := append(resJ, arr...)
				newRes = append(newRes, temp)
			}
		}
		res = append(res, newRes...)
		res = append(res, numArr...)
	}
	res = append(res, []int{})
	return res
}
