package hw23

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func getData() []struct {
	text     string
	pattern  string
	expected int
} {
	return []struct {
		text     string
		pattern  string
		expected int
	}{
		{text: "Not found", pattern: "some", expected: -1},
		{text: "Not found in text", pattern: "текст", expected: -1},
		{text: "aaaaaaaaaaaaba", pattern: "aaaaaaaaaab", expected: 2},
		{text: "aaabaaaaaaaa", pattern: "baaa", expected: 3},
		{text: "aababaaaaaaaa", pattern: "baba", expected: 2},
		{text: "bacdbaddba", pattern: "baddba", expected: 4}, // здесь нужно при смещении использовать префикс
		{text: "caddbaddba", pattern: "baddba", expected: 4}, // здесь нужно при смещении использовать префикс
		{text: "колокол", pattern: "окол", expected: 3},
		{text: "First", pattern: "Fi", expected: 0},
		{text: "Found in text", pattern: "text", expected: 9},
		{text: "Я иду гуляю по Москве", pattern: "по", expected: 12},
	}
}

func TestExhaustiveSearch(t *testing.T) {
	data := getData()
	search := ExhaustiveSearch{}
	for _, v := range data {
		result := search.Search(v.text, v.pattern)
		require.Equal(t, v.expected, result)
	}
}
