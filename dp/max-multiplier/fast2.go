package maxmultiplier

func maximumScoreF2(nums []int, muls []int) int {
	lenMul := len(muls)
	memo := make([][]int, lenMul+1)
	for ii := 0; ii <= lenMul; ii++ {
		memo[ii] = make([]int, lenMul+1)
	}
	return dpF(0, 0, nums, muls, memo)
}

func dpF(l, i int, nums []int, muls []int, memo [][]int) int {
	if i == len(muls) {
		return 0 // Picked enough m elements
	}
	if memo[l][i] != 0 {
		return memo[l][i]
	}
	pickLeft := dpF(l+1, i+1, nums, muls, memo) + nums[l]*muls[i]                // Pick the left side
	pickRight := dpF(l, i+1, nums, muls, memo) + nums[len(nums)-(i-l)-1]*muls[i] // Pick the right side
	res := max(pickLeft, pickRight)
	memo[l][i] = res
	return res
}
