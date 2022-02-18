package consecutivenumberssum

func consecutiveNumbersSum(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	count := 1
	startNbr := 1
	endNbr := 2
	acc := 3
	for {
		if acc == n {
			count++
			acc = acc - startNbr - endNbr
			startNbr++
			endNbr--
		}
		if n < acc {
			acc = acc - startNbr - endNbr
			startNbr++
			endNbr--
		}
		if n < startNbr*2-1 {
			break
		}
		endNbr++
		acc += endNbr
	}
	return count
}
