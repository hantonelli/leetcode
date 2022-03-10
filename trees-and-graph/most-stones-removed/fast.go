package most_stones_removed

func removeStonesFast(stones [][]int) int {
	rows := make(map[int][]int)
	cols := make(map[int][]int)

	for _, stone := range stones {
		rows[stone[0]] = append(rows[stone[0]], stone[1])
		cols[stone[1]] = append(cols[stone[1]], stone[0])
	}

	unreachableStones := 0
	visited := make(map[[2]int]bool)

	for _, stone := range stones {
		if _, ok := visited[[2]int{stone[0], stone[1]}]; !ok {
			unreachableStones++
			dfs(stone[0], stone[1], rows, cols, visited)
		}
	}
	return len(stones) - unreachableStones
}

func dfs(x, y int, rows, cols map[int][]int, visited map[[2]int]bool) {
	visited[[2]int{x, y}] = true
	for _, col := range rows[x] {
		if !visited[[2]int{x, col}] {
			dfs(x, col, rows, cols, visited)
		}
	}
	for _, row := range cols[y] {
		if !visited[[2]int{row, y}] {
			dfs(row, y, rows, cols, visited)
		}
	}
}
