package q16

func intersection(nums1 []int, nums2 []int) []int {
	mp := make(map[int]struct{})
	var res []int
	for _, num := range nums1 {
		mp[num] = struct{}{}
	}
	for _, num := range nums2 {
		if _, ok := mp[num]; ok {
			res = append(res, num)
			delete(mp, num)
		}
	}
	return res
}
