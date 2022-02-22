package main

import (
	"sort"
	"testing"

	dataStructures "github.com/hantonelli/leetcode/data-structures"
	"github.com/stretchr/testify/require"
)

var base []int = dataStructures.RandArr(1000000)

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
