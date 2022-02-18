package maximumsubarray

import "math"

func maxSubArrayFast(nums []int) int {
	var max_so_far int = math.MinInt32 // to track sum
	var max_ending_here int = 0        // end of contiguous array
	var l int = len(nums)

	for i := 0; i < l; i++ {
		max_ending_here = max_ending_here + nums[i]

		if max_so_far < max_ending_here {
			max_so_far = max_ending_here
		}
		// If it is negative, we can set max_ending_here to 0
		// and then the next sum will be just the next number
		if max_ending_here < 0 {
			max_ending_here = 0
		}
	}

	return max_so_far

}
