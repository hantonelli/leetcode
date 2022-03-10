package next_closest_time

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		time     string
		expected string
	}{
		// {
		// 	time:     "19:34",
		// 	expected: "19:39",
		// },
		// {
		// 	time:     "23:59",
		// 	expected: "22:22",
		// },
		// {
		// 	time:     "13:55",
		// 	expected: "15:11",
		// },
		// {
		// 	time:     "22:23",
		// 	expected: "22:32",
		// },
		// {
		// 	time:     "12:22",
		// 	expected: "21:11",
		// },
		{
			time:     "23:49",
			expected: "22:22",
		},
	}

	for _, tt := range tests {
		require.EqualValues(t, tt.expected, nextClosestTime(tt.time), tt.time)
	}
}
