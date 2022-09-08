package license_key_formatting

import "strings"

func licenseKeyFormattingNoFast(s string, k int) string {
	r := strings.NewReplacer("-", "")
	s = r.Replace(s)
	s = strings.ToUpper(s)

	parts := make([]string, 0, len(s)/k)
	firstSize := len(s) % k
	if firstSize > 0 {
		parts = append(parts, s[:firstSize])
	}
	for i := firstSize; i < len(s); i += k {
		parts = append(parts, s[i:i+k])
	}
	return strings.Join(parts, "-")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
