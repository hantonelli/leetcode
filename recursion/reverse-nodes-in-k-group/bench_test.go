package reverse_nodes_in_k_group

import (
	"testing"

	"github.com/hantonelli/leetcode/ds"
)

const (
	k         = 10000
	nodeCount = 100000000
)

var arr = ds.RandArr(nodeCount)

// func Benchmark_Single_Sol(b *testing.B) {
// 	data := ds.ArrToLinkedList(arr)
// 	b.ResetTimer()
// 	reverseKGroup(data, k)
// }

func Benchmark_Single_Fast(b *testing.B) {
	data := ds.ArrToLinkedList(arr)
	b.ResetTimer()
	reverseKGroupFast(data, k)
}

func Benchmark_Single_Small(b *testing.B) {
	data := ds.ArrToLinkedList(arr)
	b.ResetTimer()
	reverseKGroupSmall(data, k)
}
