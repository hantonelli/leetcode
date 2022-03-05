package ds

import "testing"

var base []int = RandArr(10000)

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := NewStack()
		for i := 0; i < len(base); i++ {
			s.Push(base[i])
		}
		for s.Size() != 0 {
			s.Pop()
		}
	}
}
func Benchmark_Fast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := NewStack2()
		for i := 0; i < len(base); i++ {
			s.Push(base[i])
		}
		for s.Size() != 0 {
			s.Pop()
		}
	}
}
func Benchmark_Generic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := NewStack3[int]()
		for i := 0; i < len(base); i++ {
			s.Push(base[i])
		}
		for s.Size() != 0 {
			s.Pop()
		}
	}
}
