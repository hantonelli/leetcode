package wordbreak

func wordBreak(s string, wordDict []string) bool {
	for _, w := range wordDict {
		if len(s) < len(w) {
			continue
		}
		if s[:len(w)] == w {
			if len(s) == len(w) {
				return true
			}
			if can := wordBreak(s[len(w):], wordDict); can {
				return true
			}
		}
	}
	return false
}
