package nthtribonaccinumber

func tribonacci(n int) int {
	memo := make([]int, n+1)
	return dp(n, memo)
}

func dp(n int, memo []int) int {
	if v := memo[n]; v != 0 {
		return v
	}
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}
	acc := dp(n-1, memo) + dp(n-2, memo) + dp(n-3, memo)
	memo[n] = acc
	return acc
}
