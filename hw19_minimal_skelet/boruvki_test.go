package hw19

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBoruvkilSkeletonExists(t *testing.T) {
	graph := getGraph()
	alg := NewBoruvki(graph)
	edges, err := alg.GetSkeleton()
	require.NoError(t, err)
	validEdges := []Edge{{v1: 0, v2: 1}, {v1: 2, v2: 3}, {v1: 1, v2: 2}}
	require.Equal(t, validEdges, edges)
}

// TestBoruvkiSkeletonNotExists у одной из вершины нет исходящих/входящих ребер.
func TestBoruvkiSkeletonNotExists(t *testing.T) {
	graph := getGraph()
	graph = append(graph, []*GraphEdge{nil, nil, nil})

	alg := NewBoruvki(graph)
	_, err := alg.GetSkeleton()
	require.Error(t, err)
}
