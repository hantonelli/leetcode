package meetingrooms2

import (
	"sort"

	ds "github.com/hantonelli/leetcode/data-structures"
)

func minMeetingRooms(intervals [][]int) int {
	sort.SliceStable(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	stack := [][]int{intervals[0]}
	maxStack := 1
	for i := 1; i < len(intervals); i++ {
		for j := 0; j < len(stack); j++ {
			if stack[j][1] <= intervals[i][0] {
				stack = append(stack[:j], stack[j+1:]...)
			}
		}
		stack = append(stack, intervals[i])
		maxStack = ds.Max(maxStack, len(stack))
	}
	return maxStack
}
