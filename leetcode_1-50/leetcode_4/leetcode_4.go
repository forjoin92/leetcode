package leetcode_4

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	newNums := make([]int, 0, len(nums1)+len(nums2))
	for len(nums1) != 0 || len(nums2) != 0 {
		if len(nums1) == 0 {
			newNums = append(newNums, nums2[0])
			nums2 = nums2[1:]
			continue
		}
		if len(nums2) == 0 {
			newNums = append(newNums, nums1[0])
			nums1 = nums1[1:]
			continue
		}
		if nums1[0] < nums2[0] {
			newNums = append(newNums, nums1[0])
			nums1 = nums1[1:]
		} else {
			newNums = append(newNums, nums2[0])
			nums2 = nums2[1:]
		}
	}
	median1, median2 := 0, 0
	length := len(newNums)
	if length%2 == 0 {
		median1 = newNums[length/2-1]
		median2 = newNums[length/2]
		return (float64(median1) + float64(median2)) / 2
	}
	return float64(newNums[length/2])
}
