package countofsmallernumbersafterself

import (
	"testing"

	dataStructures "github.com/hantonelli/leetcode/data-structures"
)

var base []int = dataStructures.RandArr(1000)

// func Benchmark_Sol(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		countSmaller(base)
// 	}
// }

func Benchmark_Sol2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countSmaller2(base)
	}
}

func Benchmark_Fastest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		countSmallerFastest(base)
	}
}
