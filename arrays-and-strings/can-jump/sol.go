package canjump

func canJump(nums []int) bool {
	if len(nums) < 2 {
		return true
	}
	reachable := make([]bool, len(nums))
	reachable[0] = true
	for i, v := range nums {
		if reachable[i] {
			if len(nums)-1 <= i+v {
				return true
			}
			for j := 1; j <= v; j++ {
				if !reachable[i+j] {
					reachable[i+j] = true
				}
			}
		}
	}
	return false
}
