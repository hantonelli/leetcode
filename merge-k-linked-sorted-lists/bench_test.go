package main

import (
	"testing"

	"github.com/hantonelli/leetcode/ds"
	ll "github.com/hantonelli/leetcode/ds"
)

var base []int = ds.RandArr(1000000)

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
