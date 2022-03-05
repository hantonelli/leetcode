package minimumwindowsubstring

import "github.com/hantonelli/leetcode/ds"

func minWindow(s string, t string) string {
	tMap := map[rune]int{}
	for _, r := range t {
		tMap[r]++
	}
	starts := []int{}
	sCharPos := map[rune][]int{}
	found := false
	for pos, r := range s {
		if _, tOk := tMap[r]; tOk {
			if _, sok := sCharPos[r]; !sok {
				sCharPos[r] = []int{pos}
			} else {
				sCharPos[r] = append(sCharPos[r], pos)
			}
			starts = append(starts, pos)
		}
	}

	minStart, minEnd, minLen := 0, 0, len(s)
	for _, start := range starts {
		end := 0
		possible := true
		for ch, countReq := range tMap {
			chPos, ok := sCharPos[ch]
			if !ok {
				possible = false
				break
			}

			for i := 0; i < len(chPos) && 0 < countReq; i++ {
				if start <= chPos[i] {
					countReq--
					end = ds.Max(end, chPos[i])
				}
			}
			if countReq != 0 {
				possible = false
				break
			}
		}
		if possible && (end-start) < minLen {
			found = true
			minStart = start
			minEnd = end
			minLen = end - start
		}
	}
	if !found {
		return ""
	}
	return s[minStart : minEnd+1]
}

// func minWindow(s string, t string) string {
// 	sMap := map[rune]int{}
// 	for _, r := range t {
// 		sMap[r]--
// 	}
// 	tLeft := len(t)
// 	for i:=0; i< len(s); i++ {
// 		for _, r := range s {
// 			if v, ok := sMap[r]; ok {
// 				if v < 0 {
// 					tLeft--
// 				}
// 			}
// 		}

// 	}

// 	return ""
// }
