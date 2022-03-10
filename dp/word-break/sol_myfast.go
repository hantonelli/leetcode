package wordbreak

func wordBreakMyFast(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 0; i < len(s); i++ {
		if !dp[i] {
			continue
		}
		for _, w := range wordDict {
			if i+len(w) <= len(s) && s[i:i+len(w)] == w {
				dp[i+len(w)] = true
			}
		}
	}
	return dp[len(dp)-1]
}
