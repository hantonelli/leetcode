package mincostclimbingstairs

func minCostClimbingStairs2(cost []int) int {
	if len(cost) == 0 {
		return 0
	}
	if len(cost) == 1 {
		return cost[0]
	}
	if len(cost) == 0 {
		return min(cost[0], cost[1])
	}
	dp := make([]int, len(cost))
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2; i < len(cost); i++ {
		dp[i] = cost[i] + min(dp[i-1], dp[i-2])
	}
	return min(dp[len(cost)-2], dp[len(cost)-1])
}
