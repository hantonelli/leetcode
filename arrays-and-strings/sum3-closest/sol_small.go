package sum3closest

import "sort"

func threeSumClosestSmall(nums []int, target int) int {
	sort.Ints(nums)
	sum := 0
	res := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums)-2; i++ {
		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum = nums[i] + nums[left] + nums[right]
			if sum == target {
				return sum
			}
			if abs(sum-target) < abs(res-target) {
				res = sum
			}
			if sum < target {
				left++
			} else {
				right--
			}
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
