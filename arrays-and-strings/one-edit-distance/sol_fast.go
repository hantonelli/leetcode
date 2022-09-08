package one_edit_distance

func isOneEditDistanceFast(s string, t string) bool {
	// if this question asks about edit distance, it's actually easy to code
	if len(s) > len(t) {
		s, t = t, s
	}

	// now we have len(s) <= len(t)
	if len(t)-len(s) > 1 {
		return false
	} else if len(t) == len(s) {
		i := len(t) - 1
		for i >= 0 && t[i] == s[i] {
			i--
		}
		if i == -1 {
			return false
		} else {
			return s[:i] == t[:i]
		}
	} else {
		i := 0
		for i < len(s) && t[i] == s[i] {
			i++
		}
		// we runinto t[i] != s[i], and we know t is longer than s by one
		// what we can do is remove this elements from t, see if that make the two the same
		return s[i:] == t[i+1:]
	}
}
