package main

func Fib(n int) int {
	if n <= 2 {
		return 1
	}
	memo := map[int]int{}
	return fib(n-1, memo) + fib(n-2, memo)
}

func fib(n int, memo map[int]int) int {
	if _, ok := memo[n]; ok {
		return memo[n]
	}
	if n <= 2 {
		return 1
	}
	memo[n] = fib(n-1, memo) + fib(n-2, memo)
	return memo[n]
}
