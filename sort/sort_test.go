package sort

import (
	"sort"
	"testing"

	"github.com/hantonelli/leetcode/ds"
	"github.com/stretchr/testify/require"
)

var base []int = ds.RandArr(10)

func TestMerge(t *testing.T) {
	input := make([]int, len(base))
	copy(input, base)
	expected := make([]int, len(base))
	copy(expected, input)
	sort.Ints(expected)
	require.Equal(t, expected, MergeSort(input))
}

func TestLibrary(t *testing.T) {
	input := make([]int, len(base))
	copy(input, base)
	expected := make([]int, len(base))
	copy(expected, input)
	sort.Ints(expected)
	sort.Ints(input)
	require.Equal(t, expected, input)
}

func TestQuicksort(t *testing.T) {
	input := make([]int, len(base))
	copy(input, base)
	expected := make([]int, len(base))
	copy(expected, input)
	sort.Ints(expected)
	Quicksort(input)
	require.Equal(t, expected, input)
}

func TestQuicksortFn(t *testing.T) {
	input := make([]int, len(base))
	cmp := func(a, b int) bool {
		return a < b
	}
	copy(input, base)
	expected := make([]int, len(base))
	copy(expected, input)
	sort.Ints(expected)
	QuicksortFn(input, cmp)
	require.Equal(t, expected, input)
}
