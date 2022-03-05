package meetingrooms2

import (
	"sort"

	ds "github.com/hantonelli/leetcode/data-structures"
)

func minMeetingRoomsFast(intervals [][]int) int {
	if len(intervals) < 2 {
		return 1
	}

	count, res := 0, 0
	start := make([]int, 0, len(intervals))
	end := make([]int, 0, len(intervals))

	for i := 0; i < len(intervals); i++ {
		start = append(start, intervals[i][0])
		end = append(end, intervals[i][1])
	}

	sort.Ints(start)
	sort.Ints(end)

	s := 0
	e := 0
	for e < len(intervals) && s < len(intervals) {
		if start[s] < end[e] {
			count += 1
			s++
		} else {
			count -= 1
			e++
		}
		res = ds.Max(res, count)
	}
	return res
}
