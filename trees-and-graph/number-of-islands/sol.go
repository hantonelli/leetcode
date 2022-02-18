package main

func numIslands(grid [][]byte) int {
	count := 0

	visited := map[[2]int]bool{}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' && !visited[[2]int{i, j}] {
				count++
				visited[[2]int{i, j}] = true
				neigboards := [][2]int{{i, j}}
				for 0 < len(neigboards) {
					nei := neigboards[len(neigboards)-1]
					neigboards = neigboards[:len(neigboards)-1]
					posI := nei[0]
					posJ := nei[1]

					if posI+1 < len(grid) && !visited[[2]int{posI + 1, posJ}] && grid[posI+1][posJ] == '1' {
						visited[[2]int{posI + 1, posJ}] = true
						neigboards = append(neigboards, [2]int{posI + 1, posJ})
					}
					if posJ+1 < len(grid[0]) && !visited[[2]int{posI, posJ + 1}] && grid[posI][posJ+1] == '1' {
						visited[[2]int{posI, posJ + 1}] = true
						neigboards = append(neigboards, [2]int{posI, posJ + 1})
					}
					if 0 <= posI-1 && !visited[[2]int{posI - 1, posJ}] && grid[posI-1][posJ] == '1' {
						visited[[2]int{posI - 1, posJ}] = true
						neigboards = append(neigboards, [2]int{posI - 1, posJ})
					}
					if 0 <= posJ-1 && !visited[[2]int{posI, posJ - 1}] && grid[posI][posJ-1] == '1' {
						visited[[2]int{posI, posJ - 1}] = true
						neigboards = append(neigboards, [2]int{posI, posJ - 1})
					}
				}
			}
		}
	}
	return count
}
