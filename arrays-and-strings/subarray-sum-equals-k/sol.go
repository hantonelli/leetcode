package subarray_sum_equals_k

import "sort"

func subarraySum(nums []int, k int) int {
	sort.Ints(nums)
	res := 0
	currentSum := 0
	for i := 0; i < len(nums); i++ {
		start, end := i, i
		currentSum = nums[start]
		for start < len(nums)-1 {
			if currentSum == k {
				res++
				currentSum -= nums[start]
				start++
				if end < start {
					end++
					currentSum = nums[start]
				}
			} else {
				if currentSum < k {
					end++
					if end == len(nums)-1 {
						start++
						end = start
					}
					currentSum += nums[end]
				} else {
					currentSum -= nums[start]
					start++
					if end < start {
						end++
						currentSum = nums[start]
					}
				}
			}
		}
	}
	return res
}
