package mergeintervals

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		intervals [][]int
		expected  [][]int
	}{
		{
			intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			expected:  [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			intervals: [][]int{{1, 4}, {4, 5}},
			expected:  [][]int{{1, 5}},
		},
		{
			intervals: [][]int{{1, 4}, {0, 4}},
			expected:  [][]int{{0, 4}},
		},
		{
			intervals: [][]int{{1, 4}, {0, 5}},
			expected:  [][]int{{0, 5}},
		},
	}

	for _, tt := range tests {
		require.EqualValues(t, tt.expected, merge(tt.intervals))
	}
}
