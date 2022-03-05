package findduplicatesubtrees

import (
	"testing"

	ds "github.com/hantonelli/leetcode/data-structures"
)

// var base []int = ds.RandArr(1000000)
// var root = ds.ListToTree(base)
var root = ds.ListToTree([]int{1, 2, 3, 4, ds.NULL, 2, 4, ds.NULL, ds.NULL, 4})

func Benchmark_Sol(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findDuplicateSubtrees(root)
	}
}

func Benchmark_Fast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findDuplicateSubtreesFast(root)
	}
}
