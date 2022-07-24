package hw19

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// getGraph возвращает граф в виде A[N][Smax] - несуществующие элементы добиваем nil
// можно также завести специальное значение GraphEdge и ссылаться уже на него
// Граф - взвешенный неориентированный.
func getGraph() [][]*GraphEdge {
	return [][]*GraphEdge{
		{&GraphEdge{1, 5}, &GraphEdge{2, 8}, nil},
		{&GraphEdge{0, 5}, &GraphEdge{2, 7}, &GraphEdge{3, 17}},
		{&GraphEdge{0, 8}, &GraphEdge{1, 7}, &GraphEdge{3, 6}},
		{&GraphEdge{1, 17}, nil, nil},
	}
}

func TestKraskalSkeletonExists(t *testing.T) {
	graph := getGraph()
	alg := NewKraskal(graph)
	edges, err := alg.GetSkeleton()
	require.NoError(t, err)
	validEdges := []Edge{{v1: 0, v2: 1}, {v1: 2, v2: 3}, {v1: 1, v2: 2}}
	require.Equal(t, validEdges, edges)
}

// TestKraskalSkeletonNotExists у одной из вершины нет исходящих/входящих ребер.
func TestKraskalSkeletonNotExists(t *testing.T) {
	graph := getGraph()
	graph = append(graph, []*GraphEdge{nil, nil, nil})

	alg := NewKraskal(graph)
	_, err := alg.GetSkeleton()
	require.Error(t, err)
}
