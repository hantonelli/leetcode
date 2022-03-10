package androidunlockpatterns

func numberOfPatterns2(m int, n int) int {
	var cnt int
	backtrack(m, n, 0, 0, make(map[int]bool), &cnt)
	return cnt
}

func backtrack(m, n int, lastN int, seqLength int, visited map[int]bool, cnt *int) {
	if seqLength > n {
		return
	}

	if seqLength >= m {
		(*cnt)++
	}

	for i := 1; i <= 9; i++ {
		if _, ok := visited[i]; ok {
			continue
		}

		visited[i] = true

		if seqLength == 0 || canVisit(visited, lastN, i) {
			backtrack(m, n, i, seqLength+1, visited, cnt)
		}
		delete(visited, i)
	}

}

func canVisit(visited map[int]bool, src, dst int) bool {
	srcX, srcY := getCoordinatesFromNumber(src)
	dstX, dstY := getCoordinatesFromNumber(dst)

	// allows visiting adjacent numbers on a dialpad
	if abs(dstX-srcX) <= 1 && abs(dstY-srcY) <= 1 {
		return true
	}

	// allows visiting numbers that doesn't have other number on a route
	// e.g. 1->6, 1->8
	if (abs(dstX-srcX) == 2 && abs(dstY-srcY) == 1) || (abs(dstX-srcX) == 1 && abs(dstY-srcY) == 2) {
		return true
	}

	// allows visiting number if other number in a middle is visited
	midX, midY := (srcX+dstX)/2, (srcY+dstY)/2
	mid := getNumberFromCoordinates(midX, midY)

	return visited[mid]
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func getNumberFromCoordinates(x, y int) int {
	return (y)*3 + x + 1
}

func getCoordinatesFromNumber(n int) (x, y int) {
	return (n - 1) % 3, (n - 1) / 3
}
