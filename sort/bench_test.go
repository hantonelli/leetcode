package sort

import (
	"math/rand"
	"sort"
	"testing"
	// "github.com/hantonelli/leetcode/ds"
)

func GetIntNArray(l, n int) []int {
	if l <= 0 {
		return make([]int, 0)
	}
	seed := int64(666)
	r := rand.New(rand.NewSource(seed))
	arr := make([]int, l)
	for i := 0; i < l; i++ {
		arr[i] = r.Intn(n)
	}
	return arr
}

var (
	// base2  []int = ds.RandArr(10000)
	base2  []int = GetIntNArray(10000, 20000)
	copies       = 60000000 / len(base2)
	res    []int
)

func generateBaseCopies(base []int, copies int) [][]int {
	res := make([][]int, copies)
	for i := 0; i < copies; i++ {
		cp := make([]int, len(base))
		copy(cp, base)
		res[i] = cp
	}
	return res
}

func BenchmarkLibrary(b *testing.B) {
	copies := generateBaseCopies(base2, copies)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.Ints(copies[i])
	}
}

func BenchmarkLibraryStable(b *testing.B) {
	copies := generateBaseCopies(base2, copies)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.SliceStable(copies[i], func(a, b int) bool {
			return copies[i][a] < copies[i][b]
		})
	}
}

func BenchmarkLibraryUnstable(b *testing.B) {
	copies := generateBaseCopies(base2, copies)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.Slice(copies[i], func(a, b int) bool {
			return copies[i][a] < copies[i][b]
		})
	}
}

func Benchmark_QuickSort(b *testing.B) {
	copies := generateBaseCopies(base2, copies)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Quicksort(copies[i])
	}
}

func Benchmark_QuickSortFn(b *testing.B) {
	cmp := func(a, b int) bool {
		return a < b
	}
	copies := generateBaseCopies(base2, copies)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		QuicksortFn(copies[i], cmp)
	}
}

// func BenchmarkMergeSort(b *testing.B) {
// 	copies := generateBaseCopies(base2, copies)
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		res = MergeSort(copies[i])
// 	}
// }
