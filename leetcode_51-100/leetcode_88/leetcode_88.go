package leetcode_88

func merge(nums1 []int, m int, nums2 []int, n int) {
	pos := len(nums1) - 1
	for m > 0 && n > 0 {
		if nums1[m-1] > nums2[n-1] {
			nums1[pos] = nums1[m-1]
			m--
		} else {
			nums1[pos] = nums2[n-1]
			n--
		}
		pos--
	}
	if n > 0 {
		copy(nums1, nums2[:n])
	}
}
