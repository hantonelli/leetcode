package medianoftwosortedarrays

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	leng := len(nums1) + len(nums2)
	pos := int(math.Ceil(float64(leng) / 2))
	isPair := leng%2 == 0
	var j, k, prev int
	var lastJ bool

	for i := 0; ; i++ {
		if !isPair && i == pos {
			if lastJ {
				return float64(nums1[j-1])
			} else {
				return float64(nums2[k-1])
			}
		}
		if isPair && i == pos {
			if lastJ {
				prev = nums1[j-1]
			} else {
				prev = nums2[k-1]
			}
		}
		if isPair && i == pos+1 {
			if lastJ {
				return float64(prev+nums1[j-1]) / 2
			} else {
				return float64(prev+nums2[k-1]) / 2
			}
		}
		if j == len(nums1) {
			lastJ = false
			k++
			continue
		}
		if k == len(nums2) {
			lastJ = true
			j++
			continue
		}
		if nums1[j] < nums2[k] {
			lastJ = true
			j++
		} else {
			lastJ = false
			k++
		}
	}
}
