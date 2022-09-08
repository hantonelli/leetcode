package valid_palindrome_2

func validPalindrome(s string) bool {
	i, j := 0, len(s)-1
	diffFound := false
	for i < j {
		if s[i] == s[j] {
			i++
			j--
			continue
		}
		if diffFound {
			return false
		}
		// we reach the center, and if we remove a letter, then the word it valid
		if j-i == 1 {
			break
		}
		sub1 := s[i+1 : j+1]
		sub2 := s[i:j]
		return validPalindromeStrict(sub1) || validPalindromeStrict(sub2)
	}
	return true
}

func validPalindromeStrict(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
