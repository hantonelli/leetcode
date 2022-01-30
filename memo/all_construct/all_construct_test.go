package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAllConstruct(t *testing.T) {
	require.EqualValues(t, [][]string{
		{"le", "purp"},
		{"le", "p", "ur", "p"},
	}, AllConstruct("purple", []string{"purp", "p", "ur", "le", "purpl"}))

	require.Equal(t, [][]string{
		{"ef", "cd", "ab"},
		{"def", "c", "ab"},
		{"def", "abc"},
		{"ef", "abcd"},
	}, AllConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd", "ef", "c"}))

	require.Equal(t, [][]string{}, AllConstruct("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}))

	require.Equal(t, [][]string{
		{"ot", "p", "ent", "t", "p", "a", "enter"},
		{"t", "o", "p", "ent", "ot", "p", "a", "enter"},
		{"ot", "p", "ent", "t", "o", "p", "a", "enter"},
		{"t", "o", "p", "ent", "t", "o", "p", "a", "enter"},
	}, AllConstruct("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"}))

	require.Equal(t, [][]string{}, AllConstruct("eeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{"e", "ee", "eee", "eeee", "eeeee", "eeeeee", "eeeeeee"}))
}
