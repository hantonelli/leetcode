package main

func Fib(n int) int {
	values := make(map[int]int, n+1)
	values[0] = 1

	for i := range values {
		if i+1 < len(values) {
			values[i+1] = values[i]
		}
		if i+2 < len(values) {
			values[i+2] = values[i]
		}
	}

	return values[n]
}
