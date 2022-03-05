package minimumcosttohirekworkers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		quality  []int
		wage     []int
		k        int
		expected float64
	}{
		{
			quality:  []int{10, 20, 5},
			wage:     []int{70, 50, 30},
			k:        2,
			expected: 105.00000,
		},
		{
			quality:  []int{3, 1, 10, 10, 1},
			wage:     []int{4, 8, 2, 2, 7},
			k:        3,
			expected: 30.66667,
		},
		{
			quality:  []int{32, 43, 66, 9, 94, 57, 25, 44, 99, 19},
			wage:     []int{187, 366, 117, 363, 121, 494, 348, 382, 385, 262},
			k:        4,
			expected: 1528.00000,
		},
		{
			quality:  []int{32, 43, 57, 44},
			wage:     []int{187, 366, 494, 382},
			k:        4,
			expected: 1528.00000,
		},
		{
			quality:  []int{32, 43, 66, 94, 57, 44},
			wage:     []int{187, 366, 117, 121, 494, 382},
			k:        4,
			expected: 1528.00000,
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, mincostToHireWorkers3(tt.quality, tt.wage, tt.k))
	}
}
