package countofsmallernumbersafterself

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
			input:    []int{5, 2, 6, 1},
			expected: []int{2, 1, 1, 0},
		},
		{
			input:    []int{-1},
			expected: []int{0},
		},
		{
			input:    []int{-1, -1},
			expected: []int{0, 0},
		},
		{
			input:    []int{27, 26, 26, 100},
			expected: []int{2, 0, 0, 0},
		},
	}

	for _, tt := range tests {
		require.EqualValues(t, tt.expected, countSmaller(tt.input))
	}
}
