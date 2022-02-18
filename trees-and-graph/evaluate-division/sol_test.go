package evaluateDivision

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		equations [][]string
		values    []float64
		queries   [][]string
		expected  []float64
	}{
		{
			equations: [][]string{{"a", "b"}, {"b", "c"}},
			values:    []float64{2.0, 3.0},
			queries:   [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}},
			expected:  []float64{6.00000, 0.50000, -1.00000, 1.00000, -1.00000},
		},
		{
			equations: [][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}},
			values:    []float64{1.5, 2.5, 5.0},
			queries:   [][]string{{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"}},
			expected:  []float64{3.75000, 0.40000, 5.00000, 0.20000},
		},
		{
			equations: [][]string{{"a", "b"}},
			values:    []float64{0.5},
			queries:   [][]string{{"a", "b"}, {"b", "a"}, {"a", "c"}, {"x", "y"}},
			expected:  []float64{0.50000, 2.00000, -1.00000, -1.00000},
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, calcEquation(tt.equations, tt.values, tt.queries))
	}
}
