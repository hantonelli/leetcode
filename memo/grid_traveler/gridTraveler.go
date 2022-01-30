package main

import "fmt"

func GridTraveler(m, n int) int {
	memo := map[string]int{}
	return gridTraveler(m, n, memo)
}

func gridTraveler(m, n int, memo map[string]int) int {
	key := fmt.Sprintf("%d,%d", m, n)
	if _, ok := memo[key]; ok {
		return memo[key]
	}
	if m == 1 && n == 1 {
		return 1
	}
	if m == 0 || n == 0 {
		return 0
	}

	memo[key] = gridTraveler(m-1, n, memo) + gridTraveler(m, n-1, memo)
	return memo[key]
}
