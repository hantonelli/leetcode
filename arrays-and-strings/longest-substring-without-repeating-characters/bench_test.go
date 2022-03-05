package main

import "testing"

func Benchmark_Sol(b *testing.B) {
	input := "pwwkewpkoadslparfs0irq3knAFIH21K32OKNFADS90PKODASDPOK"
	for i := 0; i < b.N; i++ {
		lengthOfLongestSubstring(input)
	}
}

func Benchmark_Fast(b *testing.B) {
	input := "pwwkewpkoadslparfs0irq3knAFIH21K32OKNFADS90PKODASDPOK"
	for i := 0; i < b.N; i++ {
		lengthOfLongestSubstringFast(input)
	}
}
