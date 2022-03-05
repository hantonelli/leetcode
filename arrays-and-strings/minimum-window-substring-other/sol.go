package minimumwindowsubstringother

func minWindow(s string, t string) string {
	if len(t) == 0 || len(s) == 0 {
		return ""
	}
	charsPos := make(map[rune]*[]int, 26*2)
	for i, r := range s {
		if _, ok := charsPos[r]; ok {
			*charsPos[r] = append(*charsPos[r], i)
		} else {
			charsPos[r] = &[]int{i}
		}
	}
	if len(t) == 1 {
		_, ok := charsPos[rune(s[0])]
		if ok {
			return t
		}
	}
	initLet, ok := charsPos[rune(t[0])]
	if !ok {
		return ""
	}
	var startPick, endPick, lenPick int
	// Loop over possible start positions (t[0]->all t[0] contained in s)
	for _, sIdx := range *initLet {
		start := sIdx
		last := start
		valid := true
		// Loop over the rest of the letters in t, and look for the closest following s for each
		for i := 1; i < len(t); i++ {
			// looking for letter t[i]

			// loop over all s position of t[i]
			sPosForTi, ok := charsPos[rune(t[i])]
			if !ok {
				valid = false
				break
			}
			found := false
			for _, trIdx := range *sPosForTi {
				if last < trIdx {
					found = true
					last = trIdx
					break
				}
			}
			if !found {
				valid = false
				break
			}
		}
		currLen := last - start
		if valid && (lenPick == 0 || currLen < lenPick) {
			startPick = start
			endPick = last
			lenPick = currLen
		}
	}
	return s[startPick : endPick+1]
}
