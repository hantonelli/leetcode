package main

import (
	"math"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"

	ll "github.com/hantonelli/leetcode/data-structures"
)

func randArr(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(math.MaxInt64)
	}
	return arr
}

var base []int = randArr(1000000)

func Test(t *testing.T) {
	tests := []struct {
		input    []*ll.ListNode
		expected *ll.ListNode
	}{
		{
			input:    ll.ArrOfArrToArrLinkedList([][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}),
			expected: ll.ArrToLinkedList([]int{1, 1, 2, 3, 4, 4, 5, 6}),
		},
	}

	for _, tt := range tests {
		res := mergeKListsF(tt.input)
		require.Equal(t, tt.expected, res)
	}
}

var f *ll.ListNode

func Benchmark(b *testing.B) {
	input := ll.ArrOfArrToArrLinkedList([][]int{base, base})
	b.ResetTimer()
	f = mergeKLists(input)
}

func BenchmarkF(b *testing.B) {
	input := ll.ArrOfArrToArrLinkedList([][]int{base, base})
	b.ResetTimer()
	f = mergeKListsF(input)
}

func BenchmarkFast(b *testing.B) {
	input := ll.ArrOfArrToArrLinkedList([][]int{base, base})
	b.ResetTimer()
	f = mergeKListsFast(input)
}
