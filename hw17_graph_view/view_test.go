package graphview

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGraphViewFirst(t *testing.T) {
	factory := NewViewFactory("./views/1.txt")

	matrix := factory.GetAdjacencyMatrix()
	ar := factory.GetAdjacencyArray()
	tb := factory.GetTableContents()

	degreeTests := []struct {
		expected int
		vertex   int
	}{
		{expected: 1, vertex: 1},
		{expected: 3, vertex: 2},
		{expected: 2, vertex: 5},
	}

	for _, val := range degreeTests {
		val := val
		t.Run(fmt.Sprintf("vertex %d", val.vertex), func(t *testing.T) {
			require.Equal(t, val.expected, ar.Degree(val.vertex))
			require.Equal(t, val.expected, matrix.Degree(val.vertex))
			require.Equal(t, val.expected, tb.Degree(val.vertex))
		})
	}

	isAdjacencyTests := []struct {
		expected bool
		v1       int
		v2       int
	}{
		{expected: true, v1: 1, v2: 2},
		{expected: true, v1: 2, v2: 1},
		{expected: false, v1: 1, v2: 4},
		{expected: true, v1: 5, v2: 5},
		{expected: false, v1: 1, v2: 1},
	}

	for _, val := range isAdjacencyTests {
		val := val
		t.Run(fmt.Sprintf("vertex %d %d", val.v1, val.v2), func(t *testing.T) {
			require.Equal(t, val.expected, ar.IsAdjacency(val.v1, val.v2))
			require.Equal(t, val.expected, matrix.IsAdjacency(val.v1, val.v2))
			require.Equal(t, val.expected, tb.IsAdjacency(val.v1, val.v2))
		})
	}

	ListAdjacencyTests := []struct {
		expected []int
		v        int
	}{
		{expected: []int{2}, v: 1},
		{expected: []int{1, 3, 4}, v: 2},
		{expected: []int{5}, v: 5},
	}

	for _, val := range ListAdjacencyTests {
		val := val
		t.Run(fmt.Sprintf("list vertex %d", val.v), func(t *testing.T) {
			list := ar.ListAdjacency(val.v)
			sort.Ints(list)
			require.Equal(t, val.expected, list)
			list = matrix.ListAdjacency(val.v)
			sort.Ints(list)
			require.Equal(t, val.expected, list)
			list = tb.ListAdjacency(val.v)
			sort.Ints(list)
			require.Equal(t, val.expected, list)
		})
	}
}
