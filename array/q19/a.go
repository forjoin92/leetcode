package q19

func validMountainArray(A []int) bool {
	if len(A) < 3 {
		return false
	}
	if A[0] >= A[1] {
		return false
	}
	trend := true
	for i := 0; i < len(A); i++ {
		if i > 0 {
			if A[i-1] == A[i] {
				return false
			}
			if A[i-1] > A[i] {
				trend = false
			} else {
				if !trend {
					return false
				}
			}
		}
	}
	if !trend {
		return true
	}
	return false
}
