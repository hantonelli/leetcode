package coinchange

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	lefts := []int{amount}
	visited := map[int]bool{}
	for deep := 1; 0 < len(lefts); deep++ {
		next := []int{}
		for _, left := range lefts {
			visited[left] = true
			for _, c := range coins {
				if c == left {
					return deep
				}
				if c < left && !visited[left-c] {
					next = append(next, left-c)
				}
			}
		}
		lefts = next
	}
	return -1
}
