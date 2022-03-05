package readncharactersgivenread4

var solutionSmaller = func(read4 func([]byte) int) func([]byte, int) int {
	// implement read below.
	return func(buf []byte, n int) int {
		copiedChars, readChars := 0, 4
		for copiedChars < n && readChars == 4 {
			readChars = read4(buf[copiedChars:])
			copiedChars += readChars
		}
		return Min(n, copiedChars)
	}
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
