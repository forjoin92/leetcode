package q4

func containsDuplicate(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}

	mp := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		if _, ok := mp[nums[i]]; ok {
			return true
		}
		mp[nums[i]] = struct{}{}
	}
	return false
}
