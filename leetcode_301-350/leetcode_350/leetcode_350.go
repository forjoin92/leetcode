package leetcode_350

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}

	m1 := make(map[int]int, len(nums1))
	m2 := make(map[int]int, len(nums2))
	for _, num := range nums1 {
		m1[num]++
	}
	for _, num := range nums2 {
		m2[num]++
	}

	var res []int
	var base, iterator map[int]int
	if len(m1) < len(m2) {
		iterator = m1
		base = m2
	} else {
		iterator = m2
		base = m1
	}
	for num, times := range iterator {
		value, exist := base[num]
		if exist {
			if times < value {
				for i := 0; i < times; i++ {
					res = append(res, num)
				}
			} else {
				for i := 0; i < value; i++ {
					res = append(res, num)
				}
			}
		}
	}
	return res
}
