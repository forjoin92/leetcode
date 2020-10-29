package q12

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	array := make([][]byte, numRows)
	f := true
	now := 0
	for _, c := range s {
		array[now] = append(array[now], byte(c))
		if f {
			if now == numRows-1 {
				now--
				f = false
			} else {
				now++
			}
		} else {
			if now == 0 {
				now++
				f = true
			} else {
				now--
			}
		}
	}
	var newArray []byte
	for i := 0; i < numRows; i++ {
		newArray = append(newArray, array[i]...)
	}
	return string(newArray)
}
