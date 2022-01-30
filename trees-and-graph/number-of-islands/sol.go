package main

func numIslands(grid [][]byte) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				count++
				sinkIsland(grid, i, j)
			}
		}
	}
	return count
}

func sinkIsland(grid [][]byte, i, j int) {
	grid[i][j] = 0
	if 0 < i-1 && grid[i-1][j] == '1' {
		sinkIsland(grid, i-1, j)
	}
	if i+1 < len(grid) && grid[i+1][j] == '1' {
		sinkIsland(grid, i+1, j)
	}
	if 0 < j-1 && grid[i][j-1] == '1' {
		sinkIsland(grid, i, j-1)
	}
	if j+1 < len(grid[i]) && grid[i][j+1] == '1' {
		sinkIsland(grid, i, j+1)
	}
}
