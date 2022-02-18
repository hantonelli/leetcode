package insertinterval

func insertFast(intervals [][]int, newInterval []int) [][]int {
	ans := [][]int{}
	for i := 0; i < len(intervals); i++ {
		// if no interval in future matches this one
		if newInterval[1] < intervals[i][0] {
			// add new interval
			ans = append(ans, newInterval[:])
			// append rest of them
			ans = append(ans, intervals[i:]...)
			return ans
		} else if newInterval[0] > intervals[i][1] {
			ans = append(ans, intervals[i])
		} else {
			newInterval[0] = min(newInterval[0], intervals[i][0])
			newInterval[1] = max(newInterval[1], intervals[i][1])
		}
	}

	// if not returned then there are extending intervals
	ans = append(ans, newInterval[:])
	return ans
}
