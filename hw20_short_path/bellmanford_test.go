package hw20

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func getGraphWithNegativeCycle() [][]*GraphEdge {
	return [][]*GraphEdge{
		// A
		{&GraphEdge{1, 4}, &GraphEdge{2, 7}},
		// B
		{&GraphEdge{2, 2}, nil},
		// C
		{&GraphEdge{3, -3}, nil},
		// D
		{&GraphEdge{0, -4}, nil},
	}
}

func getGraphWithNegativeEdges() [][]*GraphEdge {
	return [][]*GraphEdge{
		// A
		{&GraphEdge{1, 7}, &GraphEdge{2, 5}, &GraphEdge{3, 2}},
		// B
		{&GraphEdge{2, -3}, nil, nil},
		// C
		{&GraphEdge{3, -3}, nil, nil},
		// D
		{nil, nil, nil},
	}
}

func TestBellmanExistsLongPath(t *testing.T) {
	graph := getGraph()
	alg := NewBellmanFord(graph)
	edges, weight, err := alg.GetMinPath(0, 3)
	require.NotEqual(t, notSet, weight)
	require.Nil(t, err)
	// A -> B -> C -> D
	validEdges := []Edge{{v1: 0, v2: 1}, {v1: 1, v2: 2}, {v1: 2, v2: 3}}
	validWeight := 6

	require.Equal(t, validEdges, edges)
	require.Equal(t, validWeight, weight)
}

func TestBellmanExistsSamePath(t *testing.T) {
	graph := getGraph()
	alg := NewBellmanFord(graph)
	edges, weight, err := alg.GetMinPath(0, 0)
	require.NotEqual(t, notSet, weight)
	require.Nil(t, err)
	require.Nil(t, edges)
	require.Equal(t, 0, weight)
}

func TestBellmanNotExistsPath(t *testing.T) {
	graph := getGraph()
	alg := NewBellmanFord(graph)
	// C -> A
	edges, weight, err := alg.GetMinPath(2, 0)
	require.Equal(t, notSet, weight)
	require.Nil(t, err)
	require.Nil(t, edges)
}

func TestBellmanNegativeCycle(t *testing.T) {
	graph := getGraphWithNegativeCycle()
	alg := NewBellmanFord(graph)
	// C -> A
	_, _, err := alg.GetMinPath(2, 0)
	require.Error(t, err)
	require.True(t, errors.Is(err, ErrNegativeCycle))
}

func TestBellmanExistsNegativeLongPath(t *testing.T) {
	graph := getGraphWithNegativeEdges()
	alg := NewBellmanFord(graph)
	edges, weight, err := alg.GetMinPath(0, 3)
	require.NotEqual(t, notSet, weight)
	require.Nil(t, err)
	// A -> B -> C -> D
	validEdges := []Edge{{v1: 0, v2: 1}, {v1: 1, v2: 2}, {v1: 2, v2: 3}}
	validWeight := 1

	require.Equal(t, validEdges, edges)
	require.Equal(t, validWeight, weight)
}
