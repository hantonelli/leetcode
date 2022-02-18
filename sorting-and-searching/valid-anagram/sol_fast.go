package validanagram

func isAnagramFast(s string, t string) bool {
	var i int
	size1 := int('z' - 'a' + 1)
	chars := make([]byte, size1)
	for i = 0; i < len(s); i++ {
		k1 := s[i] - 'a'
		chars[k1]++
	}
	for i = 0; i < len(t); i++ {
		chars[t[i]-'a']--
	}
	for i = 0; i < len(chars); i++ {
		if chars[i] != 0 {
			return false
		}
	}
	return true
}
