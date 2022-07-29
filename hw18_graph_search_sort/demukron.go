package hw18

import (
	"errors"
)

var ErrSortNotExists = errors.New("topological sort not exists")

// Demukron - топологическая сортировка графа заданным вектором смежности int A[N][Smax]
// Нумерация вершин графа начинается с нуля
// Значений которых не существует добъем с помощью -1.
type Demukron struct {
	graph [][]int
}

func (d *Demukron) Sort() ([][]int, error) {
	var result [][]int
	var level int
	vector := make([]int, len(d.graph))

	for _, vertexes := range d.graph {
		for _, v1 := range vertexes {
			if v1 == -1 {
				break
			}
			vector[v1]++
		}
	}

	countVertex := len(d.graph)
	for countVertex > 0 {
		result = append(result, []int{})

		for i, s := range vector {
			if s == 0 {
				result[level] = append(result[level], i)
				vector[i] = -1
			}
		}

		count := len(result[level])
		if count == 0 {
			return nil, ErrSortNotExists
		}

		for _, v := range result[level] {
			for _, v1 := range d.graph[v] {
				if v1 == -1 {
					break
				}
				vector[v1]--
			}
		}
		countVertex -= count
		level++
	}
	return result, nil
}

func NewDemukron(graph [][]int) *Demukron {
	return &Demukron{
		graph: graph,
	}
}
