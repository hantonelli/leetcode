package minimumwindowsubstring

func minWindowOld(s string, t string) string {
	var r rune
	var sStartPtr, sEndPtr, minLen int
	charInT := map[rune]int{}
	for _, r := range t {
		charInT[r]++
	}
	charInS := map[rune]int{}
	for _, r := range t {
		charInS[r] = 0
	}
	charsFound := 0

	for i := 0; i < len(s); i++ {
		r = rune(s[i])
		if _, ok := charInT[r]; !ok {
			continue
		}

		charsFound = 0
		for j := i; j < len(s); j++ {
			r = rune(s[j])
			if _, ok := charInT[r]; ok {
				if charInS[r] < charInT[r] {
					charInS[r]++
					charsFound++

					if charsFound == len(t) {
						if minLen == 0 || (j-i) < minLen {
							minLen = j - i
							sStartPtr = i
							sEndPtr = j + 1
						}
						break
					}
				}
			}
			if minLen != 0 && minLen < (j-i) {
				break
			}
		}

		for _, r := range t {
			charInS[r] = 0
		}
	}

	if sEndPtr == 0 {
		return ""
	}
	return s[sStartPtr:sEndPtr]
}

// func minWindow(s string, t string) string {
// 	var sStartPtr, sEndPtr, minLen int

// 	for i := 0; i < len(s); i++ {
// 		if s[i] != t[0] {
// 			continue
// 		}
// 		tIdn := 0
// 		for j := i; j < len(s); j++ {
// 			if s[j] == t[tIdn] {
// 				tIdn++
// 			}
// 			if tIdn == len(t) {
// 				if minLen == 0 || (j-i) < minLen {
// 					minLen = j - i
// 					sStartPtr = i
// 					sEndPtr = j
// 				}
// 				break
// 			}
// 		}
// 	}
// 	if minLen == 0 {
// 		return ""
// 	}
// 	return s[sStartPtr : sEndPtr+1]
// }

// Any order, WITHOUT REPEATING T
// func minWindow(s string, t string) string {
// 	charInT := map[rune]bool{}
// 	for _, r := range t {
// 		charInT[r] = true
// 	}

// 	var start, end int
// 	length := len(s)
// 	var newStart, newEnd int
// 	charPosition := map[rune]int{}
// 	for i, r := range s {
// 		if charInT[r] {
// 			charPosition[r] = i

// 			if len(charPosition) == len(t) {
// 				newStart, newEnd = len(s)-1, 0
// 				for _, pos := range charPosition {
// 					if pos < newStart {
// 						newStart = pos
// 					}
// 					if newEnd < pos {
// 						newEnd = pos
// 					}
// 				}
// 				if newEnd-newStart < length {
// 					start = newStart
// 					end = newEnd
// 					length = newEnd - newStart
// 				}
// 			}
// 		}
// 	}

// 	return s[start : end+1]
// }
