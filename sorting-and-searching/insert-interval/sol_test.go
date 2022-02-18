package insertinterval

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		intervals   [][]int
		newInterval []int
		expected    [][]int
	}{
		{
			intervals:   [][]int{{1, 3}, {6, 9}},
			newInterval: []int{2, 5},
			expected:    [][]int{{1, 5}, {6, 9}},
		},
		{
			intervals:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			newInterval: []int{4, 8},
			expected:    [][]int{{1, 2}, {3, 10}, {12, 16}},
		},
		{
			intervals:   [][]int{{1, 10}},
			newInterval: []int{2, 4},
			expected:    [][]int{{1, 10}},
		},
		{
			intervals:   [][]int{{2, 4}},
			newInterval: []int{1, 10},
			expected:    [][]int{{1, 10}},
		},
		{
			intervals:   [][]int{},
			newInterval: []int{5, 7},
			expected:    [][]int{{5, 7}},
		},
		{
			intervals:   [][]int{{1, 5}},
			newInterval: []int{6, 8},
			expected:    [][]int{{1, 5}, {6, 8}},
		},
		{
			intervals:   [][]int{{1, 5}},
			newInterval: []int{0, 0},
			expected:    [][]int{{0, 0}, {1, 5}},
		},
		{
			intervals:   [][]int{{2, 4}, {5, 7}, {8, 10}, {11, 13}},
			newInterval: []int{3, 6},
			expected:    [][]int{{2, 7}, {8, 10}, {11, 13}},
		},
		{
			intervals:   [][]int{{3, 5}, {12, 15}},
			newInterval: []int{6, 6},
			expected:    [][]int{{3, 5}, {6, 6}, {12, 15}},
		},
	}

	for _, tt := range tests {
		require.EqualValues(t, tt.expected, insert(tt.intervals, tt.newInterval))
	}
}
