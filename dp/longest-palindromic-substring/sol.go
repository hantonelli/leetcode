package longestpalindromicsubstring

func longestPalindrome(s string) string {
	longest := ""
	longestLen := -1
	for i := range s {
		for j := len(s) - 1; i < j && (longestLen == -1 || longestLen < (j-i)); j-- {
			if s[i] == s[j] {
				if isPalindrome(s[i : j+1]) {
					if longestLen == -1 || longestLen < (j-i) {
						longestLen = j - i
						longest = s[i : j+1]
					}
				}
			}
		}
	}
	if longestLen == -1 {
		return s[:1]
	}
	return longest
}

func isPalindrome(s string) bool {
	leng := len(s) - 1
	halfLen := int(len(s) / 2)
	for i := 0; i < halfLen; i++ {
		if s[i] != s[leng-i] {
			return false
		}
	}
	return true
}
