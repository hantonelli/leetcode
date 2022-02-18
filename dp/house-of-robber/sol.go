package houseofrobber

func rob(nums []int) int {
	memo := make(map[int]int, len(nums))
	return dp(0, nums, memo)
}

func dp(i int, nums []int, memo map[int]int) int {
	if v, ok := memo[i]; ok {
		return v
	}
	if len(nums) <= i {
		return 0
	}
	curr := max(nums[i]+dp(i+2, nums, memo), dp(i+1, nums, memo))
	memo[i] = curr
	return curr
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
