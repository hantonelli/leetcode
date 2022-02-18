package mincostclimbingstairs

func minCostClimbingStairs(cost []int) int {
	memo := make(map[int]int, len(cost))
	return dp(-1, cost, memo)
}

func dp(i int, cost []int, memo map[int]int) int {
	if len(cost) <= i {
		return 0
	}
	if v, ok := memo[i]; ok {
		return v
	}
	nextCost := min(dp(i+1, cost, memo), dp(i+2, cost, memo))
	if i != -1 {
		nextCost += cost[i]
	}
	memo[i] = nextCost
	return nextCost
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
