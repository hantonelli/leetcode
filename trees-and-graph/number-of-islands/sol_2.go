package main

import "container/list"

func numIslands2(grid [][]byte) int {
	count := 0
	visited := map[int]bool{}
	neigboards := list.List{}

	lenI := len(grid)
	lenJ := len(grid[0])
	var ni, nj, pos int

	for i := 0; i < lenI; i++ {
		for j := 0; j < lenJ; j++ {
			if grid[i][j] == '1' && !visited[i*lenJ+j] {
				count++
				neigboards.PushBack(i*lenJ + j)

				for 0 < neigboards.Len() {
					el := neigboards.Front()
					pos = el.Value.(int)
					neigboards.Remove(el)
					ni = int(pos / lenJ)
					nj = pos % lenJ
					visited[ni*lenJ+nj] = true

					if 0 <= ni-1 && grid[ni-1][nj] == '1' && !visited[(ni-1)*lenJ+nj] {
						neigboards.PushBack((ni-1)*lenJ + nj)
					}
					if ni+1 < lenI && grid[ni+1][nj] == '1' && !visited[(ni+1)*lenJ+nj] {
						neigboards.PushBack((ni+1)*lenJ + nj)
					}
					if 0 <= nj-1 && grid[ni][nj-1] == '1' && !visited[ni*lenJ+nj-1] {
						neigboards.PushBack(ni*lenJ + nj - 1)
					}
					if nj+1 < lenJ && grid[ni][nj+1] == '1' && !visited[ni*lenJ+nj+1] {
						neigboards.PushBack(ni*lenJ + nj + 1)
					}
				}
			}
		}
	}
	return count
}
