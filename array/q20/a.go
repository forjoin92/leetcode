package q20

func insert(intervals [][]int, newInterval []int) [][]int {
	var res [][]int
	f := false
	if len(intervals) == 0 {
		return append(res, newInterval)
	}
	i := 0
	for i < len(intervals) {
		if newInterval[1] < intervals[i][0] {
			res = append(res, newInterval)
			f = true
			break
		} else if newInterval[1] == intervals[i][0] {
			res = append(res, []int{newInterval[0], intervals[i][1]})
			f = true
			break
		} else if newInterval[1] > intervals[i][0] {
			if newInterval[1] <= intervals[i][1] {
				if newInterval[0] >= intervals[i][0] {
					res = append(res, []int{intervals[i][0], intervals[i][1]})
				} else {
					res = append(res, []int{newInterval[0], intervals[i][1]})
				}
				f = true
				break
			} else {
				if newInterval[0] <= intervals[i][0] {
					res = append(res, []int{newInterval[0], newInterval[1]})
					f = true
					break
				} else {
					if newInterval[0] <= intervals[i][1] {
						res = append(res, []int{intervals[i][0], newInterval[1]})
						f = true
						break
					} else {
						res = append(res, intervals[i])
						i++
						continue
					}
				}
			}
		}
	}
	if !f {
		res = append(res, newInterval)
	} else {
		res = append(res, intervals[i:]...)
	}
	res = deal(res)
	return res
}

func deal(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}
	var res [][]int
	head, tail := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if tail < intervals[i][0] {
			res = append(res, []int{head, tail})
			head, tail = intervals[i][0], intervals[i][1]
		} else if tail == intervals[i][0] {
			tail = intervals[i][1]
		} else {
			if tail <= intervals[i][1] {
				tail = intervals[i][1]
			}
		}
	}
	res = append(res, []int{head, tail})
	return res
}
