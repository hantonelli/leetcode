package crackingTheSafe

import (
	"fmt"
	"math"
	"strings"
)

func crackSafe(n int, k int) string {
	targetNumVisited := int(math.Pow(float64(k), float64(n)))

	strPwd := strings.Repeat("0", n)
	sbPwd := strings.Builder{}
	sbPwd.Grow(targetNumVisited + k)
	sbPwd.WriteString(strPwd)

	visitedComb := map[string]bool{}
	visitedComb[strPwd] = true

	crackSafeAfter(&sbPwd, visitedComb, targetNumVisited, n, k)

	return sbPwd.String()
}

func crackSafeAfter(pwd *strings.Builder, visitedComb map[string]bool, targetNumVisited, n, k int) bool {
	// Base case: all n-length combinations among digits 0..k-1 are visited.
	if len(visitedComb) == targetNumVisited {
		return true
	}

	// Pick the last n-1 digits added
	lastDigits := pwd.String()[pwd.Len()-n+1:]
	// Iterate all the possibles digits available in the dictionary
	for ch := 0; ch < k; ch++ {
		chStr := fmt.Sprint(ch)
		// Append the Digit to the last n-1 substring, to get n digits
		newComb := lastDigits + chStr
		// Check that that combination had not been visited yet
		if !visitedComb[newComb] {
			// If it has not, mark it as visited
			visitedComb[newComb] = true
			// Add it to the result string
			pwd.WriteString(chStr)
			// Try to keep adding base on the new result string
			if crackSafeAfter(pwd, visitedComb, targetNumVisited, n, k) {
				return true
			}
			// If it couldn't keep adding, revert combination mark and revert result string
			delete(visitedComb, newComb)
			pwdStr := pwd.String()
			pwdStr = pwdStr[:len(pwdStr)-1]
			pwd.Reset()
			pwd.WriteString(pwdStr)
		}
	}

	// Couldn't find any unvisited combination by adding a single digit at the
	// end of the previous combination. This magic early stop is base on the
	// De Bruijn sequence that stablish that it should be possible to add one
	// digit to get an extra combination each time. Wooooooo.
	return false
}
