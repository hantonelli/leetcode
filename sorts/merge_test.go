package main

import (
	"math"
	"math/rand"
	"sort"
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

func randArr(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(math.MaxInt64)
	}
	return arr
}

var base []int = randArr(1000000)

func Benchmark(b *testing.B) {
	input := make([]int, 1000000)
	copy(input, base)
	expected := make([]int, 1000000)
	copy(expected, input)
	sort.Ints(expected)
	b.ResetTimer()
	require.Equal(b, expected, MergeSort(input))
}

func Benchmark2(b *testing.B) {
	input := make([]int, 1000000)
	copy(input, base)
	expected := make([]int, 1000000)
	copy(expected, input)
	sort.Ints(expected)
	b.ResetTimer()
	sort.Ints(input)
	require.Equal(b, expected, input)
}
