package stonesRemove

func removeStones(stones [][]int) int {
	lenI := len(stones)
	lenJ := len(stones[0])
	count := 0
	for i := 0; i < lenI; i++ {
		for j := 0; j < lenJ; j++ {
			if 0 < stones[i][j] {
				count++
			}
		}
	}
	toRemove := make(map[int]int, count)
	for i := 0; i < lenI; i++ {
		for j := 0; j < lenJ; j++ {
			if 0 < stones[i][j] {
				toRemove[i*lenJ+j] = stones[i][j]
			}
		}
	}
	return countRemoveStones(lenJ, toRemove)
}

func countRemoveStones(lenJ int, toRemove map[int]int) int {
	var validToRemove []int
	for k1, v1 := range toRemove {
		if 0 < v1 {
			i1 := k1 / lenJ
			j1 := k1 % lenJ
			for k2, v2 := range toRemove {
				if 0 < v2 {
					i2 := k2 / lenJ
					j2 := k2 % lenJ
					if (i1 != i2 && j1 == j2) || (i1 == i2 && j1 != j2) {
						validToRemove = append(validToRemove, k1)
						break
					}
				}
			}
		}
	}

	maxCount := 0
	for _, pos := range validToRemove {
		toRemove[pos]--
		curr := countRemoveStones(lenJ, toRemove)
		if maxCount < curr {
			maxCount = curr
		}
		toRemove[pos]++
	}

	return maxCount + 1
}
