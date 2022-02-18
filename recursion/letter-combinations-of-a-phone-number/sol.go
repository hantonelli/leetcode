package lettercombinationsofaphonenumber

var (
	charForDigit = map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}
)

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	var res []string
	res = append(res, charForDigit[digits[0]]...)
	return process(1, res, digits)
}

func process(pos int, prevConv []string, digits string) []string {
	if pos == len(digits) {
		return prevConv
	}
	var res []string
	for _, v := range prevConv {
		for _, r := range charForDigit[digits[pos]] {
			res = append(res, v+r)
		}
	}
	return process(pos+1, res, digits)
}
