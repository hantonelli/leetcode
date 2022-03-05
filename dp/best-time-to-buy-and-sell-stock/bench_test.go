package besttimetobuyandsellstock

import (
	"testing"

	"github.com/hantonelli/leetcode/ds"
)

var base []int = ds.RandArr(100)

func Benchmark_Sol(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maxProfit(base)
	}
}
func Benchmark_Sol2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maxProfit2(base)
	}
}
