package main

func expressiveWords(s string, words []string) int {
	if len(s) == 0 || len(words) == 0 {
		return 0
	}
	count := 0
	for _, w := range words {
		sIdx, wIdx := 0, 0
		for wIdx < len(w) && sIdx < len(s) {
			if w[wIdx] != s[sIdx] {
				break
			}
			chWCount := 1
			for wIdx+1 < len(w) && w[wIdx] == w[wIdx+1] {
				wIdx++
				chWCount++
			}
			chSCount := 1
			for sIdx+1 < len(s) && s[sIdx] == s[sIdx+1] {
				sIdx++
				chSCount++
			}
			valid := chWCount == chSCount || chWCount <= chSCount && 3 <= chSCount
			if !valid {
				break
			}
			wIdx++
			sIdx++
		}
		if sIdx == len(s) && wIdx == len(w) && w[wIdx-1] == s[sIdx-1] {
			count++
		}
	}
	return count
}
