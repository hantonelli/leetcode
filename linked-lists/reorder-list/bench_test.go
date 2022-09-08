package reorder_list

import (
	"testing"

	"github.com/hantonelli/leetcode/ds"
)

const items = 1000000

var base []int = ds.RandArr(items)

func Benchmark_Sol(b *testing.B) {
	head := ds.ArrToLinkedList(base)
	b.ResetTimer()
	reorderList(head)
}

func Benchmark_Fast(b *testing.B) {
	head := ds.ArrToLinkedList(base)
	b.ResetTimer()
	reorderListFast(head)
}
