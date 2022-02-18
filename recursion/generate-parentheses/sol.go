package generateparentheses

func generateParenthesis(n int) []string {
	var acc []string
	process(n, 0, 0, "", &acc)
	return acc
}

func process(n, opened, closed int, prev string, acc *[]string) {
	if opened == n && closed == n {
		*acc = append(*acc, prev)
		return
	}
	if opened+1 <= n {
		process(n, opened+1, closed, prev+"(", acc)
	}
	if closed+1 <= n && closed+1 <= opened {
		process(n, opened, closed+1, prev+")", acc)
	}
}
