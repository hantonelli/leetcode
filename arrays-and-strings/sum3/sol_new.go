package sum3

import "sort"

func threeSumNew(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	sort.Ints(nums)
	resMap := map[[3]int]bool{}
	for i := 0; i < len(nums)-2; i++ {
		left, right := i+1, len(nums)-1
		for left < right {
			diff := nums[i] + nums[left] + nums[right]
			if diff == 0 {
				resMap[[3]int{nums[i], nums[left], nums[right]}] = true
				right--
			}
			if 0 < diff {
				right--
			} else {
				left++
			}
		}
	}
	res := make([][]int, 0, len(resMap))
	for k := range resMap {
		res = append(res, []int{k[0], k[1], k[2]})
	}
	return res
}
