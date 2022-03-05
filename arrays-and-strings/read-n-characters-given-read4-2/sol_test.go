package readncharactersgivenread42

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var read4 = func(s string) func([]byte) int {
	read := 0
	return func(b []byte) int {
		toCopy := min(4, len(s)-read)
		copy(b, s[read:read+toCopy])
		read += toCopy
		return toCopy
	}
}

func Test1(t *testing.T) {
	sol := solution(read4("abc"))

	bts1 := make([]byte, 20)
	c1 := sol(bts1, 1)
	require.Equal(t, "a", string(bts1[:c1]))

	bts2 := make([]byte, 20)
	c2 := sol(bts2, 2)
	require.Equal(t, "bc", string(bts2[:c2]))

	bts3 := make([]byte, 20)
	c3 := sol(bts3, 1)
	require.Equal(t, "", string(bts3[:c3]))
}

func Test2(t *testing.T) {
	sol := solution(read4("abcdadsadsd"))

	bts1 := make([]byte, 20)
	c1 := sol(bts1, 1)
	require.Equal(t, "a", string(bts1[:c1]))

	bts2 := make([]byte, 20)
	c2 := sol(bts2, 2)
	require.Equal(t, "bc", string(bts2[:c2]))

	bts3 := make([]byte, 20)
	c3 := sol(bts3, 8)
	require.Equal(t, "dadsadsd", string(bts3[:c3]))
}

func Test3(t *testing.T) {
	sol := solution(read4("ab"))

	bts1 := make([]byte, 20)
	c1 := sol(bts1, 1)
	require.Equal(t, "a", string(bts1[:c1]))

	bts2 := make([]byte, 20)
	c2 := sol(bts2, 2)
	require.Equal(t, "b", string(bts2[:c2]))
}
