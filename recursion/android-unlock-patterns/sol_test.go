package androidunlockpatterns

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		min      int
		max      int
		expected int
	}{
		// {
		// 	min:      1,
		// 	max:      1,
		// 	expected: 9,
		// },
		{
			min:      1,
			max:      2,
			expected: 65,
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, numberOfPatterns(tt.min, tt.max))
	}
}
