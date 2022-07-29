package graphview

type AdjacencyMatrix struct {
	m [][]int
}

func (am *AdjacencyMatrix) setAdjacencyVertex(v1, v2 int, weight int) {
	am.m[v1][v2], am.m[v2][v1] = weight, weight
}

func (am *AdjacencyMatrix) ListAdjacency(v int) []int {
	var result []int
	for k, val := range am.m[v] {
		if val != 0 {
			result = append(result, k)
		}
	}
	return result
}

func (am *AdjacencyMatrix) IsAdjacency(v1, v2 int) bool {
	return am.m[v1][v2] != 0
}

func (am *AdjacencyMatrix) Degree(v int) int {
	i := 0
	for k, val := range am.m[v] {
		if val != 0 {
			if k == v {
				i += 2
				continue
			}
			i++
		}
	}
	return i
}

func NewAdjacencyMatrix(countVertex int) *AdjacencyMatrix {
	m := make([][]int, countVertex+1)
	for i := range m {
		m[i] = make([]int, countVertex+1)
	}
	return &AdjacencyMatrix{m: m}
}
