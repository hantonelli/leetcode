package sum3

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	sum2 := map[int][][2]int{}
	for i := 0; i < len(nums)-1; i++ {
		if i+2 < len(nums) && nums[i] == nums[i+1] && nums[i+1] == nums[i+2] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			curr, ok := sum2[nums[i]+nums[j]]
			if !ok {
				curr = [][2]int{{i, j}}
			} else {
				curr = append(curr, [2]int{i, j})
			}
			sum2[nums[i]+nums[j]] = curr
		}
	}
	var opt [3]int
	resMap := map[[3]int]bool{}
	for i := range nums {
		for _, s2 := range sum2[-nums[i]] {
			if s2[0] != i && s2[1] != i {
				if nums[i] < nums[s2[0]] {
					opt = [3]int{nums[i], nums[s2[0]], nums[s2[1]]}
				} else {
					if nums[i] < nums[s2[1]] {
						opt = [3]int{nums[s2[0]], nums[i], nums[s2[1]]}
					} else {
						opt = [3]int{nums[s2[0]], nums[s2[1]], nums[i]}
					}
				}
				if _, ok := resMap[opt]; !ok {
					resMap[opt] = true
				}
			}
		}
	}
	res := make([][]int, 0, len(resMap))
	for k, _ := range resMap {
		res = append(res, []int{k[0], k[1], k[2]})
	}
	return res
}
