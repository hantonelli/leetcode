package missing_number

import "sort"

func missingNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i != nums[i] {
			return i
		}
	}
	return len(nums)
}
