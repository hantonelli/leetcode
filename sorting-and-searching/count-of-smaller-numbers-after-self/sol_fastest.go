package countofsmallernumbersafterself

func countSmallerFastest(nums []int) []int {
	/* count right part when doing merge sort descending, so O(nlgn) */
	counts := make([]int, len(nums))
	toSort := make([]int, len(nums))

	for i := range nums {
		toSort[i] = i
	}

	mergeSort(toSort, 0, len(toSort)-1, nums, counts)

	return counts
}

func mergeSort(toSort []int, s, e int, nums []int, counts []int) []int {
	if s == e {
		return toSort[s : s+1]
	}

	mid := (s + e) / 2
	left := mergeSort(toSort, s, mid, nums, counts)
	right := mergeSort(toSort, mid+1, e, nums, counts)

	merged := make([]int, e-s+1)
	ll := len(left)
	lr := len(right)
	lm := ll + lr

	for i, j, k := 0, 0, 0; k < lm; k++ {
		if i < ll && j < lr {
			if nums[left[i]] <= nums[right[j]] {
				merged[k] = right[j]
				j++
			} else {
				/* whenever we put left element in merge, we know it is larger than remaining in the right,
				 * so count of remaining in the right should be added for this left element
				 */
				merged[k] = left[i]
				counts[left[i]] += lr - j
				i++
			}
		} else if i < ll {
			merged[k] = left[i]
			i++
		} else {
			merged[k] = right[j]
			j++
		}
	}

	return merged
}
