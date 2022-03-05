package findreplacestring

func findReplaceStringFast(s string, indices []int, sources []string, targets []string) string {
	var i int

	m := map[int][2]string{}
	for i := range indices {
		m[indices[i]] = [2]string{sources[i], targets[i]}
	}

	chars := []rune(s)
	final := []rune{}
	for i < len(s) {
		val, ok := m[i]
		if !ok {
			final = append(final, chars[i])
			i++
			continue
		}
		// search for source
		if val[0] == string(chars[i:i+len(val[0])]) {
			// found, do replace
			final = append(final, []rune(val[1])...)
			i += len(val[0])
		} else {
			final = append(final, chars[i])
			i++
		}
	}
	return string(final)
}
