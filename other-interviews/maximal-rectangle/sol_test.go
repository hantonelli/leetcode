package maximalrectangle

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		matrix   [][]byte
		expected int
	}{
		{
			matrix:   [][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}},
			expected: 6,
		},
		{
			matrix:   [][]byte{{'0'}},
			expected: 0,
		},
		{
			matrix:   [][]byte{{'1'}},
			expected: 1,
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, maximalRectangleFast(tt.matrix))
	}
}
