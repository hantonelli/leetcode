package validanagram

func isAnagram(s string, t string) bool {
	sMap := map[rune]int{}
	for _, r := range s {
		sMap[r]++
	}
	for _, r := range t {
		sMap[r]--
	}
	for _, v := range sMap {
		if v != 0 {
			return false
		}
	}
	return true
}
