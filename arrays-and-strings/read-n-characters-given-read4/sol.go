package readncharactersgivenread4

var solution = func(read4 func([]byte) int) func([]byte, int) int {
	// implement read below.
	return func(buf []byte, n int) int {
		read := 0 // chars already read
		lastRead := 4
		b := make([]byte, 4)
		for lastRead == 4 {
			lastRead = read4(b)
			if lastRead == 0 {
				return read
			}
			toCopy := min(lastRead, n-read)
			copy(buf[read:], b[:toCopy])
			read += toCopy
		}
		return read
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
