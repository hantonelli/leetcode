package mergeintervals

import "sort"

func mergeFast(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		// edge cases
		if len(intervals[i]) == 0 && len(intervals[j]) == 0 {
			return false // two empty slices - so one is not less than other i.e. false
		}
		if len(intervals[i]) == 0 || len(intervals[j]) == 0 {
			return len(intervals[i]) == 0 // empty slice listed "first" (change to != 0 to put them last)
		}

		// both slices len() > 0, so can test this now:
		return intervals[i][0] < intervals[j][0]
	})

	var result [][]int
	i := 0
	f, l := intervals[i][0], intervals[i][1]
	for i < len(intervals) {
		if intervals[i][0] <= l {
			if intervals[i][1] > l {
				l = intervals[i][1]
			}
		} else {
			result = append(result, []int{f, l})
			f, l = intervals[i][0], intervals[i][1]
		}
		i += 1
	}
	result = append(result, []int{f, l})
	return result
}
