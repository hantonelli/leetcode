package permutations

import "sort"

func factorial(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func permuteSmall(nums []int) (permutations [][]int) {
	permutations = make([][]int, 0, factorial(len(nums)))
	sort.Ints(nums)
	permutations = append(permutations, nums)
	for {
		dup := make([]int, len(nums))
		copy(dup, nums)
		nums = dup
		loIdx, hiIdx := -1, 0
		for i := len(nums) - 1; i > 0; i-- {
			if nums[i-1] < nums[i] {
				loIdx = i - 1
				break
			}
		}
		if loIdx > -1 {
			for hiIdx = loIdx + 1; hiIdx < len(nums)-1; hiIdx++ {
				if nums[hiIdx+1] <= nums[loIdx] {
					break
				}
			}
			nums[loIdx], nums[hiIdx] = nums[hiIdx], nums[loIdx]
		} else {
			return
		}
		for i, j := loIdx+1, len(nums)-1; i < j; i, j = i+1, j-1 {
			nums[i], nums[j] = nums[j], nums[i]
		}
		permutations = append(permutations, nums)
	}
	return
}
