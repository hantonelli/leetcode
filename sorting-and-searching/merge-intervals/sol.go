package mergeintervals

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) < 2 {
		return intervals
	}
	sort.SliceStable(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	var res [][]int
	lastInterval := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if lastInterval[1] < intervals[i][0] {
			res = append(res, lastInterval)
			lastInterval = intervals[i]
		} else {
			if lastInterval[1] < intervals[i][1] {
				lastInterval[1] = intervals[i][1]
			}
		}
	}
	res = append(res, lastInterval)
	return res
}
