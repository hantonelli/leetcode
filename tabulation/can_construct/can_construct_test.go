package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCanConstruct(t *testing.T) {
	require.Equal(t, true, CanConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}))
	require.Equal(t, false, CanConstruct("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}))
	require.Equal(t, true, CanConstruct("enterpotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"}))
	require.Equal(t, false, CanConstruct("eeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{"e", "ee", "eee", "eeee", "eeeee", "eeeeee", "eeeeeee"}))
}
