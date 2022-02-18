package maximumsubarray

func maxSubArray(nums []int) int {
	max, currSum, start := nums[0], 0, 0
	skip := false
	for start < len(nums) {
		currSum = nums[start]
		if max < currSum {
			max = currSum
		}
		for end := start + 1; end < len(nums); end++ {
			currSum += nums[end]
			if max < currSum {
				max = currSum
			}
			if currSum < nums[end] {
				start = end
				skip = true
				break
			}
		}
		if skip {
			skip = false
		} else {
			start++
		}
	}
	return max
}
