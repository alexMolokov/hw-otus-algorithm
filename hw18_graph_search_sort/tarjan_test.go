package hw18

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestTarjanTwoVertex Есть граф состоящий из 2 вершин (0,1) из вершины 0 в ввершину 1 есть ребро 0->1.
func TestTarjanTwoVertex(t *testing.T) {
	graph := [][]int{{1}, {-1}}
	tAlg := NewTarjan(graph)
	result, err := tAlg.Sort()
	require.NoError(t, err)
	require.Equal(t, []int{0, 1}, result)
}

// TestTarjanThreeVertex Есть граф состоящий из 3 вершин (0,1,2) 2->1->0.
func TestTarjanThreeVertex(t *testing.T) {
	graph := [][]int{{-1}, {0}, {1}}
	tAlg := NewTarjan(graph)
	result, err := tAlg.Sort()
	require.NoError(t, err)
	require.Equal(t, []int{2, 1, 0}, result)
}

// TestTarjan Есть граф состоящий из 3 вершин (0,1,2) 1->0, 1->2.
func TestTarjan(t *testing.T) {
	graph := [][]int{{-1, -1}, {0, 2}, {-1, -1}}
	tAlg := NewTarjan(graph)
	result, err := tAlg.Sort()
	require.NoError(t, err)
	require.Equal(t, []int{1, 0, 2}, result)
}

// TestTarjanCycle Есть граф состоящий из 3 вершин (0,1,2) 2->1->0->1.
func TestTarjanCycle(t *testing.T) {
	graph := [][]int{{1}, {0}, {1}}
	tAlg := NewTarjan(graph)
	_, err := tAlg.Sort()
	require.Error(t, err, "not exists topological sort")
}
