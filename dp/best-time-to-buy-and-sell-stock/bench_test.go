package besttimetobuyandsellstock

import (
	"testing"

	dataStructures "github.com/hantonelli/leetcode/data-structures"
)

var base []int = dataStructures.RandArr(100)

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
