package validanagram

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	require.Equal(t, true, isAnagramFast("zlap", "kcqx"))
}
