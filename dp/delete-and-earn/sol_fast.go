package deleteandearn

func deleteAndEarn3(nums []int) int {
	mapNums := make(map[int]int)
	for _, n := range nums {
		mapNums[n]++
	}

	keys := make([]int, len(mapNums))
	i := 0
	for k := range mapNums {
		keys[i] = k
		i++
	}
	if len(keys) == 1 {
		return mapNums[keys[0]] * keys[0]
	}
	sortKeys(keys)
	memo := make([]int, len(keys))

	return max(deleteAndEarnImpl(mapNums, memo, keys, len(keys)-1), deleteAndEarnImpl(mapNums, memo, keys, len(keys)-2))
}

func deleteAndEarnImpl(mapNums map[int]int, memo []int, keys []int, x int) int {
	if x < 0 {
		return 0
	}
	if x == 0 {
		return mapNums[keys[x]] * keys[x]
	}

	if memo[x] != 0 {
		return memo[x]
	}

	next := x - 1
	if _, left := mapNums[keys[x]-1]; left {
		next--
	}

	memo[x] = mapNums[keys[x]]*keys[x] + max(deleteAndEarnImpl(mapNums, memo, keys, next), deleteAndEarnImpl(mapNums, memo, keys, next-1))
	return memo[x]
}

func sortKeys(vec []int) {
	if len(vec) < 2 {
		return
	}
	if len(vec) == 2 {
		if vec[0] > vec[1] {
			vec[0], vec[1] = vec[1], vec[0]
		}
		return
	}
	comp := vec[0]
	compIdx := 0
	idxStart := 0
	idxEnd := len(vec) - 1
	i := 0
	for i <= idxEnd {
		if comp > vec[i] {
			vec[i], vec[idxStart] = vec[idxStart], vec[i]
			idxStart++
			compIdx++
		}
		if comp < vec[i] {
			vec[i], vec[idxEnd] = vec[idxEnd], vec[i]
			idxEnd--
		} else {
			i++
		}
	}
	sortKeys(vec[:compIdx])
	if compIdx < len(vec) {
		sortKeys(vec[compIdx+1:])
	}
}
