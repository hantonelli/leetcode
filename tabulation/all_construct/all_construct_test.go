package main

// func TestAllConstruct(t *testing.T) {
// 	var nilVal [][]string
// 	require.EqualValues(t, [][]string{
// 		{"purp", "le"},
// 		{"p", "ur", "p", "le"},
// 	}, AllConstruct("purple", []string{"purp", "p", "ur", "le", "purpl"}))

// 	require.Equal(t, [][]string{
// 		{"abc", "def"},
// 		{"ab", "c", "def"},
// 		{"abcd", "ef"},
// 		{"ab", "cd", "ef"},
// 	}, AllConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd", "ef", "c"}))

// 	require.Equal(t, nilVal, AllConstruct("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}))

// 	// require.Equal(t, [][]string{
// 	// 	{"ot", "p", "ent", "t", "p", "a", "enter"},
// 	// 	{"t", "o", "p", "ent", "ot", "p", "a", "enter"},
// 	// 	{"ot", "p", "ent", "t", "o", "p", "a", "enter"},
// 	// 	{"t", "o", "p", "ent", "t", "o", "p", "a", "enter"},
// 	// }, AllConstruct("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"}))

// 	require.Equal(t, [][]string{}, AllConstruct("eeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{"e", "ee", "eee", "eeee", "eeeee", "eeeeee", "eeeeeee"}))
// }
