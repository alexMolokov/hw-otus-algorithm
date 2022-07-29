package graphview

import (
	"io/ioutil"
	"runtime"
	"strconv"
	"strings"
)

type ViewOperation interface {
	IsAdjacency(v1, v2 int) bool
	ListAdjacency(v int) []int
	Degree(v int) int
}

type edge struct {
	v1     int
	v2     int
	weight int
}

type ViewFactory struct {
	countVertex int
	countEdge   int
	edges       []edge
}

func (v *ViewFactory) GetAdjacencyMatrix() *AdjacencyMatrix {
	m := NewAdjacencyMatrix(v.countVertex)
	for _, val := range v.edges {
		m.setAdjacencyVertex(val.v1, val.v2, val.weight)
	}
	return m
}

func (v *ViewFactory) GetAdjacencyArray() *AdjacencyArray {
	m := NewAdjacencyArray(v.countVertex)
	for _, val := range v.edges {
		m.setAdjacencyVertex(val.v1, val.v2, val.weight)
	}
	return m
}

func (v *ViewFactory) GetTableContents() *TableContents {
	m := NewTableContents(v.countVertex)
	m.setAdjacencyVertex(v.edges)
	return m
}

func NewViewFactory(fileName string) *ViewFactory {
	sep := "\n"

	if runtime.GOOS == "windows" {
		sep = "\r\n"
	}

	dataIn, _ := ioutil.ReadFile(fileName)
	rows := strings.Split(string(dataIn), sep)
	info := strings.Split(rows[0], " ")

	countVertex, _ := strconv.Atoi(info[0])
	countEdge, _ := strconv.Atoi(info[1])
	edges := make([]edge, countEdge)

	lenRows := len(rows)
	for i := 1; i < lenRows; i++ {
		info := strings.Split(rows[i], " ")
		v1, _ := strconv.Atoi(info[0])
		v2, _ := strconv.Atoi(info[1])
		weight := 1
		if len(info) > 2 {
			weight, _ = strconv.Atoi(info[1])
		}
		edges[i-1] = edge{v1: v1, v2: v2, weight: weight}
	}

	return &ViewFactory{
		countVertex: countVertex,
		countEdge:   countEdge,
		edges:       edges,
	}
}
