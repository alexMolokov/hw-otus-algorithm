package hw20

import "errors"

var ErrNegativeCycle = errors.New("negative cycle. no answer")

// BellmanFord алгоритм вычисления кратчайшего пути в графе с отрицательными ребрами.
type BellmanFord struct {
	graph    [][]*GraphEdge
	distance []int
	prev     []int
}

func (b *BellmanFord) init(vertex int) {
	for i := range b.distance {
		b.distance[i] = notSet
	}
	b.distance[vertex] = 0
	for i := range b.prev {
		b.prev[i] = notSet
	}
}

// GetMinPath возвращает минимальный путь (набор ребер и значение кратчайшего пути) из v1->v2.
// Возвращает ошибку если есть цикл отрицательного веса.
func (b *BellmanFord) GetMinPath(v1, v2 int) ([]Edge, int, error) {
	b.init(v1)
	for i := len(b.distance); i >= 1; i-- {
		isUpdated := false

		for vStart := range b.graph {
			if b.distance[vStart] == notSet {
				continue
			}

			for _, edge := range b.graph[vStart] {
				if edge == nil {
					break
				}

				weight := b.distance[vStart] + edge.weight
				if b.distance[edge.v1] == notSet || b.distance[edge.v1] > weight {
					b.distance[edge.v1] = weight
					b.prev[edge.v1] = vStart
					isUpdated = true
				}
			}
		}

		if !isUpdated {
			edges, length := b.makeAnswer(v1, v2)
			return edges, length, nil
		}
	}

	return nil, 0, ErrNegativeCycle
}

func (b *BellmanFord) makeAnswer(v1, v2 int) ([]Edge, int) {
	var result []Edge
	if b.distance[v2] == notSet {
		return nil, notSet
	}

	current := v2
	for current != v1 {
		result = append(result, Edge{
			v1: b.prev[current],
			v2: current,
		})
		current = b.prev[current]
	}

	lenResult := len(result)
	for i := lenResult/2 - 1; i >= 0; i-- {
		result[i], result[lenResult-i-1] = result[lenResult-i-1], result[i]
	}
	return result, b.distance[v2]
}

func NewBellmanFord(graph [][]*GraphEdge) *BellmanFord {
	return &BellmanFord{
		graph:    graph,
		distance: make([]int, len(graph)),
		prev:     make([]int, len(graph)),
	}
}
