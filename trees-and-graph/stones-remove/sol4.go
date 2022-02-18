package stonesRemove

func removeStones4(stones [][]int) int {
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

	grid := make(map[[2]int]bool, len(stones))
	stonesInRow := make([]int, lenI)
	stonesInColumn := make([]int, lenJ)
	for _, v := range stones {
		i := v[0]
		j := v[1]
		grid[[2]int{i, j}] = true
		stonesInRow[i]++
		stonesInColumn[j]++
	}
	return countRemoveStones4(grid, stonesInRow, stonesInColumn)
}

func countRemoveStones4(stones map[[2]int]bool, stonesInRow []int, stonesInColumn []int) int {
	maxCount, i, j, curr := 0, 0, 0, 0
	for k, v := range stones {
		if v {
			i = k[0]
			j = k[1]
			if 1 < stonesInRow[i] || 1 < stonesInColumn[j] {
				stonesInRow[i]--
				stonesInColumn[j]--
				stones[[2]int{i, j}] = false
				curr = countRemoveStones4(stones, stonesInRow, stonesInColumn)
				if maxCount < curr+1 {
					maxCount = curr + 1
				}
				stonesInRow[i]++
				stonesInColumn[j]++
				stones[[2]int{i, j}] = true
			}
		}
	}
	return maxCount
}
