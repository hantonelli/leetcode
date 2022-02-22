package besttimetobuyandsellstock

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		prices   []int
		expected int
	}{
		{
			prices:   []int{7, 1, 5, 3, 6, 4},
			expected: 5,
		},
		{
			prices:   []int{7, 6, 4, 3, 1},
			expected: 0,
		},
		{
			prices:   []int{},
			expected: 0,
		},
		{
			prices:   []int{10},
			expected: 0,
		},
		{
			prices:   []int{10, 1},
			expected: 0,
		},
		{
			prices:   []int{1, 10},
			expected: 9,
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, maxProfit2(tt.prices))
	}
}
