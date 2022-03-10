package most_common_word

import (
	"strings"
)

func mostCommonWord(s string, banned []string) string {
	s = strings.ToLower(s)
	wFrec := map[string]int{}
	bannedMap := map[string]bool{}
	for _, w := range banned {
		bannedMap[w] = true
	}

	lastStart := 0
	for i := 0; i < len(s); i++ {
		if isLetter(s[i]) {
			continue
		} else {
			if i == lastStart {
				lastStart++
				continue
			}
			w := s[lastStart:i]
			if !bannedMap[w] {
				wFrec[w]++
			}
			lastStart = i + 1
		}
	}
	if lastStart != len(s) {
		w := s[lastStart:]
		if !bannedMap[w] {
			wFrec[w]++
		}
	}

	var maxFrecW string
	var maxFrecCount int
	for k, v := range wFrec {
		if maxFrecCount < v {
			maxFrecCount = v
			maxFrecW = k
		}
	}
	return maxFrecW
}

func isLetter(c byte) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z')
}
