package coinchange

import "sort"

func coinChange2(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	sort.Sort(sort.Reverse(sort.IntSlice(coins)))
	memo := map[int]int{}
	return dp(amount, coins, memo)
}

func dp(left int, coins []int, memo map[int]int) int {
	if v, ok := memo[left]; ok {
		return v
	}
	minPath := -1
	for _, c := range coins {
		if c == left {
			return 1
		}
		if c < left {
			res := dp(left-c, coins, memo)
			if res != -1 && (minPath == -1 || res+1 < minPath) {
				minPath = res + 1
			}
		}
	}
	memo[left] = minPath
	return minPath
}
