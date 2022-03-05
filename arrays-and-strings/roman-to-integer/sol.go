package romantointeger

var values = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	if len(s) == 0 {
		return 0
	}
	lastValue := 0
	total := 0
	for i := len(s) - 1; 0 <= i; i-- {
		curr, ok := values[s[i]]
		if !ok {
			panic("invalid roman value")
		}
		if lastValue <= curr {
			total += curr
		} else {
			total -= curr
		}
		lastValue = curr
	}
	return total
}
