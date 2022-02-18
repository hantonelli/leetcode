package medianoftwosortedarrays

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	tests := []struct {
		nums1    []int
		nums2    []int
		expected float64
	}{
		{
			nums1:    []int{1, 3},
			nums2:    []int{2},
			expected: 2,
		},
		{
			nums1:    []int{2},
			nums2:    []int{1, 3},
			expected: 2,
		},
		{
			nums1:    []int{1, 2},
			nums2:    []int{3, 4},
			expected: 2.5,
		},
		{
			nums1:    []int{1, 3},
			nums2:    []int{2, 4},
			expected: 2.5,
		},
		{
			nums1:    []int{1, 4},
			nums2:    []int{2, 3},
			expected: 2.5,
		},
	}

	for _, tt := range tests {
		require.EqualValues(t, tt.expected, findMedianSortedArrays(tt.nums1, tt.nums2))
	}
}
