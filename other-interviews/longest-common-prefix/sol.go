package longestcommonprefix

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	lastEq := findMinCommon(strs[0], strs[1])
	for i := 2; i < len(strs); i++ {
		if len(lastEq) == 0 {
			break
		}
		for len(lastEq) != 0 && (len(strs[i]) < len(lastEq) ||
			strs[i][:len(lastEq)] != lastEq) {
			lastEq = lastEq[:len(lastEq)-1]
		}
	}
	return lastEq
}

func findMinCommon(a, b string) string {
	eqSubstring := 0
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			break
		}
		eqSubstring = i + 1
	}
	return a[:eqSubstring]
}
