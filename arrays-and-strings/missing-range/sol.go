package missingrange

import "fmt"

func findMissingRanges(nums []int, lower int, upper int) []string {
	if len(nums) == 0 {
		if lower == upper {
			return []string{fmt.Sprintf("%d", lower)}
		} else {
			return []string{fmt.Sprintf("%d->%d", lower, upper)}
		}
	}
	res := []string{}
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			addInterval(lower-1, nums[0], &res)
		}
		if i < len(nums)-1 {
			addInterval(nums[i], nums[i+1], &res)
		} else if i == len(nums)-1 {
			addInterval(nums[i], upper+1, &res)
		}
	}
	return res
}

func addInterval(start, end int, res *[]string) {
	if start+1 <= end-1 {
		if start+1 <= end-2 {
			*res = append(*res, fmt.Sprintf("%d->%d", start+1, end-1))
		} else {
			*res = append(*res, fmt.Sprint(start+1))
		}
	}
}
