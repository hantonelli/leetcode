package one_edit_distance

func isOneEditDistance(s string, t string) bool {
	// make s the longest
	if len(s) < len(t) {
		s, t = t, s
	}
	lenDiff := len(s) - len(t)
	if len(s) == 0 || 1 < lenDiff {
		return false
	}
	diffs := 0
	j := 0
	for i := 0; i < len(s); i++ {
		if len(t) <= j {
			diffs++
			continue
		}
		if s[i] != t[j] {
			if diffs == 1 {
				return false
			}
			diffs++
			if 0 < lenDiff {
				j--
			}
		}
		j++
	}
	return diffs == 1
}
