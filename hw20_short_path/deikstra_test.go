package hw20

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// getGraph возвращает граф в виде A[N][Smax] - несуществующие элементы добиваем nil
// можно также завести специальное значение GraphEdge и ссылаться уже на него
// Граф - взвешенный ориентированный.
func getGraph() [][]*GraphEdge {
	return [][]*GraphEdge{
		// A
		{&GraphEdge{1, 3}, &GraphEdge{2, 6}, &GraphEdge{3, 9}},
		// B
		{&GraphEdge{2, 1}, nil, nil},
		// C
		{&GraphEdge{3, 2}, nil, nil},
		// D
		{nil, nil, nil},
	}
}

func TestDeikstraExistsLongPath(t *testing.T) {
	graph := getGraph()
	alg := NewDeikstra(graph)
	edges, weight := alg.GetMinPath(0, 3)
	require.NotEqual(t, notSet, weight)
	// A -> B -> C -> D
	validEdges := []Edge{{v1: 0, v2: 1}, {v1: 1, v2: 2}, {v1: 2, v2: 3}}
	validWeight := 6

	require.Equal(t, validEdges, edges)
	require.Equal(t, validWeight, weight)
}

func TestDeikstraExistsSamePath(t *testing.T) {
	graph := getGraph()
	alg := NewDeikstra(graph)
	edges, weight := alg.GetMinPath(0, 0)
	require.NotEqual(t, notSet, weight)
	require.Nil(t, edges)
	require.Equal(t, 0, weight)
}

func TestDeikstraNotExistsPath(t *testing.T) {
	graph := getGraph()
	alg := NewDeikstra(graph)
	// C -> A
	edges, weight := alg.GetMinPath(2, 0)
	require.Equal(t, notSet, weight)
	require.Nil(t, edges)
}
