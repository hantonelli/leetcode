package subsets

import "fmt"

func subsets(nums []int) [][]int {
	memo := map[string]bool{}
	perm := permute(nums)
	res := make([][]int, 0, len(perm)+1)
	res = append(res, []int{})
	for _, v := range perm {
		if _, ok := memo[fmt.Sprintf("%v", v)]; !ok {
			res = append(res, v)
			memo[fmt.Sprintf("%v", v)] = true
		}
	}
	return res
}

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{nums}
	}
	res := [][]int{nums}
	for i := 0; i < len(nums); i++ {
		res = append(res, permute(append(append([]int{}, nums[:i]...), nums[i+1:]...))...)
	}
	return res
}
