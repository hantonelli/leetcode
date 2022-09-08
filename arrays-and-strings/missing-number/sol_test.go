package missing_number

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{
			nums:     []int{3, 0, 1},
			expected: 2,
		},
		// {
		// 	nums:     []int{0, 1},
		// 	expected: 2,
		// },
		// {
		// 	nums:     []int{9, 6, 4, 2, 3, 5, 7, 0, 1},
		// 	expected: 8,
		// },
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, missingNumberFast(tt.nums))
	}
}
