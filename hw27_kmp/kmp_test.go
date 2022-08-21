package hw27

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPiBruteForce(t *testing.T) {
	tests := []struct {
		pattern  string
		expected []int
	}{
		{pattern: "AA", expected: []int{0, 0, 1}},
		{pattern: "ABC", expected: []int{0, 0, 0, 0}},
		{pattern: "ABCAB", expected: []int{0, 0, 0, 0, 1, 2}},
		{pattern: "aaBaaBaaaBa", expected: []int{0, 0, 1, 0, 1, 2, 3, 4, 5, 2, 3, 4}},
		{pattern: "kolokol", expected: []int{0, 0, 0, 0, 0, 1, 2, 3}},
		{pattern: "aabaabaaaabaabaaab", expected: []int{0, 0, 1, 0, 1, 2, 3, 4, 5, 2, 2, 3, 4, 5, 6, 7, 8, 9, 3}},
	}

	for _, v := range tests {
		require.Equal(t, v.expected, PiBruteForce(v.pattern))
	}
}

func TestPi(t *testing.T) {
	tests := []struct {
		pattern  string
		expected []int
	}{
		{pattern: "AA", expected: []int{0, 0, 1}},
		{pattern: "ABC", expected: []int{0, 0, 0, 0}},
		{pattern: "ABCAB", expected: []int{0, 0, 0, 0, 1, 2}},
		{pattern: "aaBaaBaaaBa", expected: []int{0, 0, 1, 0, 1, 2, 3, 4, 5, 2, 3, 4}},
		{pattern: "kolokol", expected: []int{0, 0, 0, 0, 0, 1, 2, 3}},
		{pattern: "aabaabaaaabaabaaab", expected: []int{0, 0, 1, 0, 1, 2, 3, 4, 5, 2, 2, 3, 4, 5, 6, 7, 8, 9, 3}},
	}

	for _, v := range tests {
		require.Equal(t, v.expected, Pi(v.pattern))
	}
}

func TestKmp(t *testing.T) {
	tests := []struct {
		text     string
		pattern  string
		expected int
	}{
		{text: "ABC", pattern: "ABC", expected: 0},
		{text: "ABRAKADABRA", pattern: "KADABRA", expected: 4},
		{text: "ABABCABD ABCDABC", pattern: "ABCD", expected: 9},
	}

	for _, v := range tests {
		require.Equal(t, v.expected, Kmp(v.text, v.pattern))
	}
}
