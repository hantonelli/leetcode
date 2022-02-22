package sum2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		input    []int
		target   int
		expected []int
	}{
		{
			input:    []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			input:    []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		{
			input:    []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
	}

	for _, tt := range tests {
		require.EqualValues(t, tt.expected, twoSum2(tt.input, tt.target))
	}
}
