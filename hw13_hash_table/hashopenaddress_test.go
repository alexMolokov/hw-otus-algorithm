package hw13hashtable

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHashOpenAddress(t *testing.T) {
	var str StringHashCode = "abc"
	hashTable := NewHashOpenAddress()
	err := hashTable.Put(&str, 1)
	require.Nil(t, err)
	require.Equal(t, 1, hashTable.len)

	var str2 StringHashCode = "acb"
	err = hashTable.Put(&str2, 2)
	require.Nil(t, err)
	require.Equal(t, 2, hashTable.len)

	result, ok := hashTable.Get(&str)
	require.True(t, ok)
	require.Equal(t, 1, result)

	result, ok = hashTable.Get(&str2)
	require.True(t, ok)
	require.Equal(t, 2, result)

	var str3 StringHashCode = "ppp"
	result, ok = hashTable.Get(&str3)
	require.False(t, ok)
	require.Nil(t, result)

	result, ok = hashTable.Remove(&str2)
	require.True(t, ok)
	require.Equal(t, 2, result)
	require.Equal(t, 1, hashTable.len)
	result, ok = hashTable.Get(&str2)
	require.False(t, ok)
	require.Nil(t, result)
}

func TestRehashOpenAddress(t *testing.T) {
	str := make([]StringHashCode, 101)
	hashTable := NewHashOpenAddress()
	for i := 100; i >= 0; i-- {
		str[i] = StringHashCode(fmt.Sprintf("%s%d", "abc", i))
		err := hashTable.Put(&str[i], i)
		require.Nil(t, err)
	}
	for i := 100; i >= 0; i-- {
		val, ok := hashTable.Get(&str[i])
		require.True(t, ok)
		require.Equal(t, i, val)
	}

	for i := 50; i >= 0; i-- {
		val, ok := hashTable.Remove(&str[i])
		require.True(t, ok)
		require.Equal(t, i, val)
	}

	for i := 50; i >= 0; i-- {
		val, ok := hashTable.Get(&str[i])
		require.False(t, ok)
		require.Nil(t, val)
	}

	for i := 100; i >= 51; i-- {
		val, ok := hashTable.Get(&str[i])
		require.True(t, ok)
		require.Equal(t, i, val)
	}

	for i := 50; i >= 0; i-- {
		err := hashTable.Put(&str[i], i)
		require.Nil(t, err)
	}

	for i := 100; i >= 0; i-- {
		val, ok := hashTable.Get(&str[i])
		require.True(t, ok)
		require.Equal(t, i, val)
	}
}
