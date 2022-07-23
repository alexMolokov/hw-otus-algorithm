package hw18

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestDemukronTwoVertex Есть граф состоящий из 2 вершин (0,1) из вершины 0 в ввершину 1 есть ребро 0->1.
func TestDemukronTwoVertex(t *testing.T) {
	graph := [][]int{{1}, {-1}}
	tAlg := NewDemukron(graph)
	result, err := tAlg.Sort()
	require.NoError(t, err)
	require.Equal(t, [][]int{{0}, {1}}, result)
}

// TestDemukronThreeVertex Есть граф состоящий из 3 вершин (0,1,2) 2->1->0.
func TestDemukronThreeVertex(t *testing.T) {
	graph := [][]int{{-1}, {0}, {1}}
	tAlg := NewDemukron(graph)
	result, err := tAlg.Sort()
	require.NoError(t, err)
	require.Equal(t, [][]int{{2}, {1}, {0}}, result)
}

// TestDemukron Есть граф состоящий из 3 вершин (0,1,2) 1->0, 1->2.
func TestDemukron(t *testing.T) {
	graph := [][]int{{-1, -1}, {0, 2}, {-1, -1}}
	tAlg := NewDemukron(graph)
	result, err := tAlg.Sort()
	require.NoError(t, err)
	require.Equal(t, [][]int{{1}, {0, 2}}, result)
}

// TestDemukronCycle Есть граф состоящий из 3 вершин (0,1,2) 2->1->0->1.
func TestDemukronCycle(t *testing.T) {
	graph := [][]int{{1}, {0}, {1}}
	tAlg := NewDemukron(graph)
	_, err := tAlg.Sort()
	require.Error(t, err, "not exists topological sort")
}
