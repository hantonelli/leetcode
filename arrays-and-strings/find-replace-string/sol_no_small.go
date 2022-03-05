package findreplacestring

import (
	"bytes"
	"sort"
)

type replacement struct {
	index  int
	source string
	target string
}

func findReplaceStringSmall(s string, indices []int, sources []string, targets []string) string {
	replacementSlc := make([]replacement, len(indices))

	for i := range indices {
		replacementSlc[i].index = indices[i]
		replacementSlc[i].source = sources[i]
		replacementSlc[i].target = targets[i]
	}

	sort.Slice(replacementSlc, func(a, b int) bool {
		return replacementSlc[a].index < replacementSlc[b].index
	})

	var currIndex, lastIndex int
	var finalStr bytes.Buffer

	for i := range replacementSlc {
		currIndex = replacementSlc[i].index
		finalStr.WriteString(s[lastIndex:currIndex])
		lastIndex = currIndex
		if s[currIndex:currIndex+len(replacementSlc[i].source)] == replacementSlc[i].source {
			finalStr.WriteString(replacementSlc[i].target)
			lastIndex += len(replacementSlc[i].source)
		}
	}

	finalStr.WriteString(s[lastIndex:len(s)])

	return finalStr.String()
}
