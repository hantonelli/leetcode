package next_closest_time

import (
	"fmt"
	"sort"
	"strconv"
)

func nextClosestTime(time string) string {
	h1, _ := strconv.Atoi(string(time[0]))
	h2, _ := strconv.Atoi(string(time[1]))
	h3, _ := strconv.Atoi(string(time[3]))
	h4, _ := strconv.Atoi(string(time[4]))
	nums := []int{h1, h2, h3, h4}
	sort.Ints(nums)
	for _, n := range nums {
		if h4 < n {
			return fmt.Sprintf("%d%d:%d%d", h1, h2, h3, n)
		}
	}
	for _, n := range nums {
		if h3 < n && (n <= 5 || (n == 6 && h4 == 0)) {
			return fmt.Sprintf("%d%d:%d%d", h1, h2, n, nums[0])
		}
	}
	for _, n := range nums {
		if h2 < n && (h1 <= 1 || (h1 == 2 && n <= 3)) {
			return fmt.Sprintf("%d%d:%d%d", h1, n, nums[0], nums[0])
		}
	}
	for _, n := range nums {
		if h1 < n && (n <= 2 || (n == 2 && h2 <= 4)) {
			return fmt.Sprintf("%d%d:%d%d", n, nums[0], nums[0], nums[0])
		}
	}
	return fmt.Sprintf("%d%d:%d%d", nums[0], nums[0], nums[0], nums[0])
}
