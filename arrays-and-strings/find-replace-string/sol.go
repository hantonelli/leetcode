package findreplacestring

import "strings"

func findReplaceString(s string, indices []int, sources []string, targets []string) string {
	sLen := len(s)
	builder := strings.Builder{}
	for i := 0; i < len(s); {
		replaced := false
		for j := 0; j < len(indices); j++ {
			if indices[j] == i && i+len(sources[j]) <= sLen && sources[j] == s[i:i+len(sources[j])] {
				replaced = true
				builder.WriteString(targets[j])
				i += len(sources[j])
				break
			}
		}
		if !replaced {
			builder.WriteByte(s[i])
			i++
		}
	}
	return builder.String()
}
