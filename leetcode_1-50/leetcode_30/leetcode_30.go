package leetcode_30

func findSubstring(s string, words []string) []int {
	if len(words) == 0 || len(words[0]) > len(s) {
		return []int{}
	}
	var res []int
	wordMap := make(map[string]bool)
	for _, word := range words {
		wordMap[word] = true
	}
	l := len(words[0])
	sArr := make([]bool, len(s)-l)
	for i := 0; i < len(s)-l; i++ {
		if _, ok := wordMap[s[i:i+l]]; ok {
			sArr[i] = true
		} else {
			sArr[i] = false
		}
	}
	for i := 0; i < len(s)-l*len(words)+1; i++ {
		start, end, match := i, i+l*len(words), true
		mp := makeWordsMap(words)
		for start < end-l+1 {
			word := s[start : start+l]
			if mp[word] > 0 {
				start += l
				mp[word]--
				continue
			}
			match = false
			break
		}
		if match {
			res = append(res, i)
		}
	}
	return res
}

func makeWordsMap(words []string) map[string]int {
	wordsMap := make(map[string]int)
	for _, word := range words {
		wordsMap[word]++
	}
	return wordsMap
}
