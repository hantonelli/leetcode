package sort

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{9, 1, 4, 2, 5, 7, 7, 9, 0},
			expected: []int{0, 1, 2, 4, 5, 7, 7, 9, 9},
		},
		{
			input:    []int{9},
			expected: []int{9},
		},
		{
			input:    []int{},
			expected: []int{},
		},
		{
			input:    []int{9, 1},
			expected: []int{1, 9},
		},
		{
			input:    []int{9, 1, 2},
			expected: []int{1, 2, 9},
		},
	}

	for _, tt := range tests {
		require.EqualValues(t, tt.expected, MergeSort(tt.input))
	}
}
