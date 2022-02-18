package main

func IsValid(s string) bool {
	stack := []rune{}
	for _, ch := range s {
		if ch == '{' || ch == '(' || ch == '[' {
			stack = append(stack, ch)
		}
		if ch == '}' || ch == ')' || ch == ']' {
			if len(stack) == 0 {
				return false
			}
			if (ch == '}' && stack[len(stack)-1] != '{') ||
				(ch == ']' && stack[len(stack)-1] != '[') ||
				(ch == ')' && stack[len(stack)-1] != '(') {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
