package groupanagrams

func groupAnagramsMyFast(strs []string) [][]string {
	keyMap := map[string][]string{}
	for i := 0; i < len(strs); i++ {
		// s := []byte(strs[i])
		// sort.Slice(s, func(i int, j int) bool { return s[i] < s[j] })
		s := sortStr(strs[i])

		if _, ok := keyMap[string(s)]; ok {
			keyMap[string(s)] = append(keyMap[string(s)], strs[i])
		} else {
			keyMap[string(s)] = []string{strs[i]}
		}
	}
	res := [][]string{}
	for _, v := range keyMap {
		res = append(res, v)
	}
	return res
}
