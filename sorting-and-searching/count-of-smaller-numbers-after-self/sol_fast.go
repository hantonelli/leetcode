package countofsmallernumbersafterself

import "sort"

func countSmallerF(nums []int) []int {
	if len(nums) < 2 {
		return []int{0}
	}
	values := map[int]int{}
	for _, n := range nums {
		values[n]++
	}
	keys := make([]int, 0, len(values))
	for k := range values {
		keys = append(keys, k)
	}
	sort.Ints(keys)

}
