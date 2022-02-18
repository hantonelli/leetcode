package consecutivenumberssum

func consecutiveNumbersSum2(n int) int {
	count := 0
	for start := 1; start <= n; start++ {
		if canGenerate2(start, n) {
			count++
		}
	}
	return count
}

func canGenerate2(start, wanted int) bool {
	acc := 0
	for start <= wanted && acc <= wanted {
		acc += start
		if acc == wanted {
			return true
		}
		start++
	}
	return false
}
