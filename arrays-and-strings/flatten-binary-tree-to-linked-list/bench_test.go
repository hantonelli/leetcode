package flatten_binary_tree_to_linked_list

import (
	"testing"

	"github.com/hantonelli/leetcode/ds"
)

var base []int = ds.RandArr(10000000)

func Benchmark(b *testing.B) {
	root := ds.ListToTree(base)
	b.ResetTimer()
	flatten(root)
}

func Benchmark_Fast(b *testing.B) {
	root := ds.ListToTree(base)
	b.ResetTimer()
	flattenFast(root)
}
