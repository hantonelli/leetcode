package most_stones_removed

func removeStones2(stones [][]int) int {
	lenI := len(stones)
	lenJ := len(stones[0])
	stonesInRow := make([]int, lenI)
	stonesInColumn := make([]int, lenJ)
	for i := 0; i < lenI; i++ {
		for j := 0; j < lenJ; j++ {
			if 0 < stones[i][j] {
				stonesInRow[i] += stones[i][j]
				stonesInColumn[j] += stones[i][j]
			}
		}
	}
	return countRemoveStones2(stones, stonesInRow, stonesInColumn)
}

func countRemoveStones2(stones [][]int, stonesInRow []int, stonesInColumn []int) int {
	maxCount := 0
	for i := 0; i < len(stones); i++ {
		for j := 0; j < len(stones[0]); j++ {
			if 0 < stones[i][j] {
				if 1 < stonesInRow[i] && 1 < stonesInColumn[j] {
					stonesInRow[i]--
					stonesInColumn[j]--
					stones[i][j]--
					curr := countRemoveStones2(stones, stonesInRow, stonesInColumn)
					if maxCount < curr {
						maxCount = curr
					}
					stonesInRow[i]++
					stonesInColumn[j]++
					stones[i][j]++
				}
			}
		}
	}
	return maxCount + 1
}
