package meetingrooms2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		intervals [][]int
		expected  int
	}{
		{
			intervals: [][]int{{0, 30}, {5, 10}, {15, 20}},
			expected:  2,
		},
		{
			intervals: [][]int{{7, 10}, {2, 4}},
			expected:  1,
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, minMeetingRooms((tt.intervals)))
	}
}
