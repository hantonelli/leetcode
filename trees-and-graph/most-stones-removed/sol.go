package most_stones_removed

func removeStones(stones [][]int) int {
	sInCol := map[int][]int{}
	sInRow := map[int][]int{}
	for _, s := range stones {
		sInCol[s[0]] = append(sInCol[s[0]], s[1])
		sInRow[s[1]] = append(sInRow[s[1]], s[0])
	}

	unreachable := 0
	visited := map[[2]int]bool{}
	for _, s := range stones {
		if _, ok := visited[[2]int{s[0], s[1]}]; !ok {
			unreachable++
			dp(s[0], s[1], sInCol, sInRow, visited)
		}
	}
	return len(stones) - unreachable
}

func dp(col, row int, sInCol, sInRow map[int][]int, visited map[[2]int]bool) {
	visited[[2]int{col, row}] = true
	for _, nextRow := range sInCol[col] {
		if !visited[[2]int{col, nextRow}] {
			dp(col, nextRow, sInCol, sInRow, visited)
		}
	}
	for _, nextCol := range sInRow[row] {
		if !visited[[2]int{nextCol, row}] {
			dp(nextCol, row, sInCol, sInRow, visited)
		}
	}
}
