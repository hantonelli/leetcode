package wordsquares

func wordSquares(words []string) [][]string {
	res := [][]string{}
	for _, w := range words {
		dpRes := dp([]string{w}, words)
		if len(dpRes) != 0 {
			res = append(res, dpRes...)
		}
	}
	return res
}

func dp(pickedWords []string, words []string) [][]string {
	if len(words[0]) == len(pickedWords) {
		return [][]string{pickedWords}
	}
	var res [][]string
	var picked string
	for _, w := range words {
		include := true
		for i := 0; i < len(pickedWords); i++ {
			picked = pickedWords[i]
			if w[i] != picked[len(pickedWords)] {
				include = false
				break
			}
		}
		if include {
			dpRes := dp(append(append(make([]string, 0, len(pickedWords)+1), pickedWords...), w), words)
			if len(dpRes) != 0 {
				res = append(res, dpRes...)
			}
		}
	}
	return res
}
