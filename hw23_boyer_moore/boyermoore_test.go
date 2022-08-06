package hw23

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBoyerMooreSearch(t *testing.T) {
	data := getData()
	search := BoyerMoore{}
	for _, v := range data {
		result := search.Search(v.text, v.pattern)
		require.Equal(t, v.expected, result)
	}
}
