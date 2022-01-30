package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		numCourses    int
		prerequisites [][]int
		expected      []int
	}{
		{
			numCourses:    2,
			prerequisites: [][]int{{1, 0}},
			expected:      []int{0, 1},
		},
		{
			numCourses:    4,
			prerequisites: [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
			expected:      []int{0, 1, 2, 3},
		},
		{
			numCourses:    1,
			prerequisites: [][]int{},
			expected:      []int{0},
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.expected, findOrder3(tt.numCourses, tt.prerequisites))
	}
}
