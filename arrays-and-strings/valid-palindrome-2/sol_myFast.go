package valid_palindrome_2

func validPalindromeMyFast(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] == s[j] {
			i++
			j--
			continue
		}
		// we reach the center, and if we remove a letter, then the word it valid
		if j-i == 1 {
			return true
		}
		validSub1 := true
		i2, j2 := i+1, j
		for i2 < j2 {
			if s[i2] != s[j2] {
				validSub1 = false
			}
			i2++
			j2--
		}
		validSub2 := true
		i2, j2 = i, j-1
		for i2 < j2 {
			if s[i2] != s[j2] {
				validSub2 = false
			}
			i2++
			j2--
		}
		return validSub1 || validSub2
	}
	return true
}
