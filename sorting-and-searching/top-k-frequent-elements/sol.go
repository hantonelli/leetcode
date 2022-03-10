package top_k_frequent_elements

import "sort"

func topKFrequent(nums []int, k int) []int {
	valFrec := map[int]int{}
	for _, val := range nums {
		valFrec[val]++
	}
	frecMap := map[int][]int{}
	for val, valFrec := range valFrec {
		if _, ok := frecMap[valFrec]; ok {
			frecMap[valFrec] = append(frecMap[valFrec], val)
		} else {
			frecMap[valFrec] = append([]int{}, val)
		}
	}
	keys := make([]int, 0, len(frecMap))
	for k2 := range frecMap {
		keys = append(keys, k2)
	}

	res := []int{}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	for i, added := 0, 0; i < len(keys) && added < k; i++ {
		for _, v := range frecMap[keys[i]] {
			added++
			res = append(res, v)
		}
	}
	return res
}
