package q23

import (
	"sort"
)

func sortByBits(arr []int) []int {
	if len(arr) == 0 {
		return nil
	}
	sort.Ints(arr)
	max := 0
	mp := make(map[int][]int)
	for _, num := range arr {
		s := 0
		tmp := num
		for tmp > 0 {
			if tmp&1 == 1 {
				s++
			}
			tmp = tmp >> 1
		}
		mp[s] = append(mp[s], num)
		if max < s {
			max = s
		}
	}
	var res []int
	for i := 0; i <= max; i++ {
		res = append(res, mp[i]...)
	}
	return res
}
