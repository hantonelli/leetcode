package loggerratelimiter

import (
	"testing"
)

func Benchmark_Sol(b *testing.B) {
	for i := 0; i < b.N; i++ {
		logger := Constructor()
		for i := 0; i < 10000000; i += 3 {
			if i%3 == 0 {
				logger.ShouldPrintMessage(i, "foo")
			}
			if i%3 == 1 {
				logger.ShouldPrintMessage(i, "bar")
			}
			if i%3 == 2 {
				logger.ShouldPrintMessage(i, "foo")
			}
		}
		logger.Close()
	}
}

func Benchmark_Small(b *testing.B) {
	for i := 0; i < b.N; i++ {
		logger := ConstructorFast()
		for i := 0; i < 10000000; i += 3 {
			if i%3 == 0 {
				logger.ShouldPrintMessage(i, "foo")
			}
			if i%3 == 1 {
				logger.ShouldPrintMessage(i, "bar")
			}
			if i%3 == 2 {
				logger.ShouldPrintMessage(i, "foo")
			}
		}
	}
}
