package maxmultiplier

func maximumScoreF1(nums []int, multipliers []int) int {
	var lenNum int = len(nums)
	var lenMul int = len(multipliers)
	if lenNum < lenMul {
		return -1
	}

	dp := make([][]int, lenMul+1)
	for ii := 0; ii <= lenMul; ii++ {
		dp[ii] = make([]int, lenMul+1)
	}

	var right, mult int
	for i := lenMul - 1; i >= 0; i-- {
		for left := i; left >= 0; left-- {
			mult = multipliers[i]
			right = lenNum - 1 - (i - left)

			dp[i][left] = max(mult*nums[left]+dp[i+1][left+1], mult*nums[right]+dp[i+1][left])
		}

	}
	return dp[0][0]
}

func peakIndexInMountainArray(arr []int) int {
	return 0
}
