package q14

func uniqueOccurrences(arr []int) bool {
	countMp := make(map[int]int)
	for _, a := range arr {
		countMp[a]++
	}
	numMp := make(map[int]struct{})
	for _, count := range countMp {
		if _, ok := numMp[count]; ok {
			return false
		}
		numMp[count] = struct{}{}
	}
	return true
}
