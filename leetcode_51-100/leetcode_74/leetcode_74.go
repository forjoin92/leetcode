package leetcode_74

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	mStart, mEnd := 0, m-1
	nStart, nEnd := 0, n-1
	if matrix[m-1][n-1] < target || matrix[0][0] > target {
		return false
	}
	for mStart <= mEnd {
		mMid := mStart + (mEnd-mStart)/2
		if matrix[mMid][n-1] == target {
			return true
		} else if matrix[mMid][n-1] < target {
			mStart = mMid + 1
		} else {
			mEnd = mMid
			for i := mEnd - 1; i >= mStart; i-- {
				if matrix[i][n-1] < target {
					mEnd = i + 1
					break
				}
				mEnd = i
			}
			if mEnd < 0 {
				mEnd = 0
			}
			break
		}
	}
	for nStart <= nEnd {
		nMid := nStart + (nEnd-nStart)/2
		if matrix[mEnd][nMid] == target {
			return true
		} else if matrix[mEnd][nMid] < target {
			nStart = nMid + 1
		} else {
			nEnd = nMid - 1
		}
	}
	return false
}
