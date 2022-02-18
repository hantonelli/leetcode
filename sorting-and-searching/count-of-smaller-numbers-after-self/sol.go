package countofsmallernumbersafterself

func countSmaller(nums []int) []int {
	if len(nums) < 2 {
		return []int{0}
	}
	values := map[int]int{}
	for _, n := range nums {
		values[n]++
	}
	var count, n, v, c int
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		n = nums[i]
		values[n]--
		count = 0
		for v, c = range values {
			if v < n && 0 < c {
				count += c
			}
		}
		res[i] = count
	}
	return res
}
