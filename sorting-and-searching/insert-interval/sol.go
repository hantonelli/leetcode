package insertinterval

func insert(intervals [][]int, newInterval []int) [][]int {
	var res [][]int
	var updated, added bool
	for i := 0; i < len(intervals); i++ {
		if intervals[i][1] < newInterval[0] || newInterval[1] < intervals[i][0] {
			if !added && (updated || newInterval[1] < intervals[i][0]) {
				updated = false
				added = true
				res = append(res, newInterval)
			}
			res = append(res, intervals[i])
			continue
		}
		updated = true
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
	}
	if !added {
		res = append(res, newInterval)
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
