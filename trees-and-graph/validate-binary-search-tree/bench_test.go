package validatebinarysearchtree

import (
	"testing"

	"github.com/hantonelli/leetcode/ds"
)

var tree = ds.ListToTree([]int{5, 4, 8, ds.NULL, ds.NULL, 7, 9, 6})

func Benchmark_Sol(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isValidBST(tree)
	}
}

func Benchmark_Fast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isValidBSTFast(tree)
	}
}
