package sum3

import (
	"testing"

	"github.com/hantonelli/leetcode/ds"
)

var base []int = ds.RandArr(100)

func Benchmark_Sol(b *testing.B) {
	for i := 0; i < b.N; i++ {
		threeSumNew(base)
	}
}

// func Benchmark_Fast(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		threeSumFast(base)
// 	}
// }
