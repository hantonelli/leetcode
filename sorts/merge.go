package main

func MergeSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	middle := len(data) / 2
	left := MergeSort(data[:middle])
	right := MergeSort(data[middle:])

	return merge(left, right)
}

func merge(l, r []int) []int {
	a := make([]int, len(l)+len(r))
	i, lIdx, rIdx := 0, 0, 0
	for lIdx < len(l) && rIdx < len(r) {
		if l[lIdx] < r[rIdx] {
			a[i] = l[lIdx]
			lIdx++
			i++
		} else {
			a[i] = r[rIdx]
			rIdx++
			i++
		}
	}
	for lIdx < len(l) {
		a[i] = l[lIdx]
		lIdx++
		i++
	}
	for rIdx < len(r) {
		a[i] = r[rIdx]
		rIdx++
		i++
	}
	return a
}
