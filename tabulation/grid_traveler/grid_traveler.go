package main

func GridTraveler(m, n int) int {
	grid := make([][]int, m+1)
	for i := range grid {
		grid[i] = make([]int, n+1)
	}
	grid[1][1] = 1

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i+1 < len(grid) {
				grid[i+1][j] += grid[i][j]
			}
			if j+1 < len(grid[i]) {
				grid[i][j+1] += grid[i][j]
			}
		}
	}
	return grid[m][n]
}
