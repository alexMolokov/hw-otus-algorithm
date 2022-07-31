package hw23

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExhaustiveSearch(t *testing.T) {
	data := []struct {
		text     string
		pattern  string
		expected int
	}{
		{text: "Not found", pattern: "some", expected: -1},
		{text: "Not found in text", pattern: "текст", expected: -1},
		{text: "First", pattern: "Fi", expected: 0},
		{text: "Found in text", pattern: "text", expected: 9},
		{text: "Я иду гуляю по Москве", pattern: "по", expected: 12},
	}

	search := ExhaustiveSearch{}
	for _, v := range data {
		result := search.Search(v.text, v.pattern)
		require.Equal(t, v.expected, result)
	}
}
