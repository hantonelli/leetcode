package main

func backspaceCompare(s string, t string) bool {
	sIdx, tIdx, toSkip := len(s)-1, len(t)-1, 0
	for 0 <= sIdx || 0 <= tIdx {
		toSkip = 0
		for 0 <= sIdx {
			if s[sIdx] == '#' {
				toSkip++
				sIdx--
			} else {
				if 0 < toSkip {
					toSkip--
					sIdx--
				} else {
					break
				}
			}
		}
		toSkip = 0
		for 0 <= tIdx {
			if t[tIdx] == '#' {
				toSkip++
				tIdx--
			} else {
				if 0 < toSkip {
					toSkip--
					tIdx--
				} else {
					break
				}
			}
		}

		if sIdx < 0 || tIdx < 0 {
			break
		}
		if s[sIdx] != t[tIdx] {
			return false
		}
		sIdx--
		tIdx--
	}

	return sIdx < 0 && tIdx < 0
}
