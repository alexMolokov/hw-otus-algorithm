package hw13hashtable

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHashCode(t *testing.T) {
	var str StringHashCode = "abc"
	code := str.GetHashCode()
	code2 := str.GetHashCode()
	require.NotEqual(t, 0, code)
	require.Equal(t, code, code2)

	var strOther StringHashCode = "acb"
	code3 := strOther.GetHashCode()

	require.NotEqual(t, code, code3)
}
