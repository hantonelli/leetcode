package lrucache

import (
	"testing"

	"github.com/hantonelli/leetcode/ds"
)

var base []int = ds.RandArr(1000)

func Benchmark_Sol(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lRUCache := Constructor(500)
		for i := 0; i < 500; i++ {
			lRUCache.Put(base[i], base[i])
		}
		for i := 500; i < 1000; i++ {
			lRUCache.Get(base[i])
		}
		for i := 0; i < 500; i++ {
			lRUCache.Get(base[i])
		}
	}
}

func Benchmark_Small(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lRUCache := Constructor(500)
		for i := 0; i < 500; i++ {
			lRUCache.Put(base[i], base[i])
		}
		for i := 500; i < 1000; i++ {
			lRUCache.Get(base[i])
		}
		for i := 0; i < 500; i++ {
			lRUCache.Get(base[i])
		}
	}
}
