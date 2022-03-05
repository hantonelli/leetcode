package readncharactersgivenread42

var solution = func(read4 func([]byte) int) func([]byte, int) int {
	b := make([]byte, 4)
	finish := false
	rest := make([]byte, 0, 4)

	// implement read below.
	return func(buf []byte, n int) int {
		read := 0 // chars already read in this invocation
		if 0 < len(rest) {
			if n <= len(rest) {
				copy(buf, rest[:n])
				rest = rest[n:]
				return n
			} else {
				copy(buf, rest)
				read += len(rest)
				rest = rest[:0]
			}
		}
		if finish {
			return read
		}
		lastRead := 4
		for !finish {
			lastRead = read4(b)
			if lastRead < 4 {
				finish = true
			}
			if lastRead == 0 {
				return read
			}
			toCopy := min(lastRead, n-read)
			copy(buf[read:], b[:toCopy])
			read += toCopy
			if lastRead != toCopy {
				rest = b[toCopy:lastRead]
				break
			}
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
