package most_stones_removed

func removeStones3(stones [][]int) int {
	var lenI, lenJ int
	for _, v := range stones {
		if lenI < v[0] {
			lenI = v[0]
		}
		if lenJ < v[1] {
			lenJ = v[1]
		}
	}
	lenI++
	lenJ++

	grid := make([][]bool, lenI)
	for i := 0; i < lenI; i++ {
		grid[i] = make([]bool, lenJ)
	}
	stonesInRow := make([]int, lenI)
	stonesInColumn := make([]int, lenJ)
	for _, v := range stones {
		i := v[0]
		j := v[1]
		grid[i][j] = true
		stonesInRow[i]++
		stonesInColumn[j]++
	}
	return countRemoveStones3(grid, stonesInRow, stonesInColumn)
}

func countRemoveStones3(stones [][]bool, stonesInRow []int, stonesInColumn []int) int {
	maxCount := 0
	for i := 0; i < len(stones); i++ {
		for j := 0; j < len(stones[0]); j++ {
			if stones[i][j] {
				if 1 < stonesInRow[i] || 1 < stonesInColumn[j] {
					stonesInRow[i]--
					stonesInColumn[j]--
					stones[i][j] = false
					curr := countRemoveStones3(stones, stonesInRow, stonesInColumn)
					if maxCount < curr+1 {
						maxCount = curr + 1
					}
					stonesInRow[i]++
					stonesInColumn[j]++
					stones[i][j] = true
				}
			}
		}
	}
	return maxCount
}
