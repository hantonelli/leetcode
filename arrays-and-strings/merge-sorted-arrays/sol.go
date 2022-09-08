package merge_sorted_arrays

func merge(nums1 []int, m int, nums2 []int, n int) {
	i1, i2, resI := m-1, n-1, len(nums1)-1
	for 0 <= resI {
		if 0 <= i1 && 0 <= i2 {
			if nums1[i1] < nums2[i2] {
				nums1[resI] = nums2[i2]
				i2--
			} else {
				nums1[resI] = nums1[i1]
				i1--
			}
		} else {
			if 0 <= i1 {
				nums1[resI] = nums1[i1]
				i1--
			} else {
				nums1[resI] = nums2[i2]
				i2--
			}
		}
		resI--
	}
}
