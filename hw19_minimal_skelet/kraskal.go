package hw19

import (
	"errors"
	"sort"

	"github.com/golang-collections/collections/stack"
)

var ErrMinimalSkeletonNotExists = errors.New("минимального остовного дерева не существует")

type Edge struct {
	v1 int
	v2 int
}

type WeightEdge struct {
	v1     int
	v2     int
	weight int
}

type GraphEdge struct {
	v1     int
	weight int
}

type Kraskal struct {
	graph      [][]*GraphEdge
	startGraph [][]int
}

// GetSkeleton возвращает минимальное остовное дерево (ребра не дублируются)
// т.е если есть ребро из вершины 0 в 1, то оно записывается только 1 раз в виде (0,1) а (1,0) опускается.
// Если дерева не существует то возвращается ошибка.
func (k *Kraskal) GetSkeleton() ([]Edge, error) {
	if k.hasZeroVertex() {
		return nil, ErrMinimalSkeletonNotExists
	}

	sortEdges := k.getSortEdges()
	result := make([]Edge, 0, len(sortEdges))

	for _, wEdge := range sortEdges {
		if k.hasPath(wEdge.v1, wEdge.v2) {
			continue
		}

		k.startGraph[wEdge.v1] = append(k.startGraph[wEdge.v1], wEdge.v2)
		k.startGraph[wEdge.v2] = append(k.startGraph[wEdge.v2], wEdge.v1)
		result = append(result, Edge{v1: wEdge.v1, v2: wEdge.v2})
	}
	return result, nil
}

// hasPath проверяет существует ли уже путь из вершины v1 в вершину v2.
func (k *Kraskal) hasPath(v1, v2 int) bool {
	used := make(map[int]bool)
	// Сознательно использую стек из библиотеки так как дз со стеком еще не делал
	st := stack.New()
	st.Push(v1)
	for st.Len() > 0 {
		el := st.Pop()
		vertex := el.(int)
		used[vertex] = true
		for _, v := range k.startGraph[vertex] {
			if _, ok := used[v]; ok {
				continue
			}
			if v == v2 {
				return true
			}
			st.Push(v)
		}
	}

	return false
}

// hasZeroVertex проверяет существуют ли вершины не соединенные с другими.
func (k *Kraskal) hasZeroVertex() bool {
	for i := range k.graph {
		if len(k.graph[i]) == 0 || k.graph[i][0] == nil {
			return true
		}
	}
	return false
}

// getSortEdges возвращает ребра в отсортированном ввиде по весу
// Также удаляет дубли ребер так как граф неориентированный.
func (k *Kraskal) getSortEdges() []*WeightEdge {
	var sortEdges []*WeightEdge
	for i := range k.graph {
		for _, v := range k.graph[i] {
			if v == nil {
				break
			}

			// ребро учтено на предыдущем шаге
			if v.v1 < i {
				continue
			}

			sortEdges = append(sortEdges, &WeightEdge{
				v1:     i,
				v2:     v.v1,
				weight: v.weight,
			})
		}
	}
	sort.Slice(sortEdges, func(i, j int) bool {
		return sortEdges[i].weight <= sortEdges[j].weight
	})
	return sortEdges
}

func NewKraskal(graph [][]*GraphEdge) *Kraskal {
	return &Kraskal{
		graph:      graph,
		startGraph: make([][]int, len(graph)),
	}
}
