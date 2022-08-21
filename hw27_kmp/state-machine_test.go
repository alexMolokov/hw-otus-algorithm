package hw27

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStateMachine(t *testing.T) {
	tests := []struct {
		text     string
		pattern  string
		expected int
	}{
		{text: "ABAC", pattern: "ABAC", expected: 0},
		{text: "ABRAKADABRA", pattern: "KADABRA", expected: 4},
		{text: "ABABCABD ABCDABC", pattern: "ABCD", expected: 9},
	}

	for _, v := range tests {
		sm := StateMachine{}
		require.Equal(t, v.expected, sm.Search(v.text, v.pattern))
	}
}
