package sum3closest

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	closestSum := math.MaxInt64
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		sum := dp(i, nums, target)
		if math.Abs(float64(target-sum)) < math.Abs(float64(target-closestSum)) {
			closestSum = sum
		}
	}
	return closestSum
}

func dp(i int, nums []int, target int) int {
	closestSum := math.MaxInt64
	j, k := 0, len(nums)-1
	for j < k {
		if j == i {
			j++
			continue
		}
		if k == i {
			k--
			continue
		}
		sum := nums[i] + nums[j] + nums[k]
		diff := target - sum
		if diff == 0 {
			return sum
		}
		if math.Abs(float64(diff)) < math.Abs(float64(target-closestSum)) {
			closestSum = sum
		}
		for j+1 < len(nums) && nums[j] == nums[j+1] && j+1 != i {
			j++
		}
		for 0 <= k-1 && nums[k] == nums[k-1] && k-1 != i {
			k--
		}
		if diff < 0 {
			k--
		} else {
			j++
		}
	}
	return closestSum
}
