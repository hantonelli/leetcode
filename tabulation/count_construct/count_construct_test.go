package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCountConstruct(t *testing.T) {
	require.Equal(t, 2, CountConstruct("purple", []string{"purp", "p", "ur", "le", "purpl"}))
	require.Equal(t, 1, CountConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}))
	require.Equal(t, 0, CountConstruct("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}))
	require.Equal(t, 4, CountConstruct("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"}))
	require.Equal(t, 0, CountConstruct("eeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{"e", "ee", "eee", "eeee", "eeeee", "eeeeee", "eeeeeee"}))
}
