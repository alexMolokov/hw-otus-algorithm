package hw16

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type myInt int

func TestBor(t *testing.T) {
	keys := []string{"my", "myKey", "manifest", "call"}
	values := []myInt{1, 2, 3, 4}
	bor := NewBor()
	for i, key := range keys {
		bor.Insert(key, values[i])
	}

	for i, key := range keys {
		require.True(t, bor.Search(key))
		val, err := bor.Get(key)
		require.NoError(t, err)
		require.Equal(t, values[i], val)
		require.True(t, bor.StartWith(key))
	}

	notExistKeys := []string{"fake", "false"}
	for _, key := range notExistKeys {
		require.False(t, bor.Search(key))
		val, err := bor.Get(key)
		require.Error(t, err)
		require.Equal(t, nil, val)
		require.False(t, bor.StartWith(key))
	}

	startWith := []string{"m", "man", "cal"}
	for _, key := range startWith {
		require.False(t, bor.Search(key))
		val, err := bor.Get(key)
		require.Error(t, err)
		require.Equal(t, nil, val)
		require.True(t, bor.StartWith(key))
	}
}
