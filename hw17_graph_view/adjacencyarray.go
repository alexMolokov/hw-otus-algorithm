package graphview

type AdjacencyVertex struct {
	v      int
	weight int
}

type AdjacencyArray struct {
	m [][]*AdjacencyVertex
}

func (am *AdjacencyArray) setAdjacencyVertex(v1, v2 int, weight int) {
	am.m[v1] = append(am.m[v1], &AdjacencyVertex{v: v2, weight: weight})
	if v1 == v2 {
		return
	}
	am.m[v2] = append(am.m[v2], &AdjacencyVertex{v: v1, weight: weight})
}

func (am *AdjacencyArray) ListAdjacency(v int) []int {
	result := make([]int, 0, len(am.m[v]))
	for _, val := range am.m[v] {
		result = append(result, val.v)
	}
	return result
}

func (am *AdjacencyArray) IsAdjacency(v1, v2 int) bool {
	for _, val := range am.m[v1] {
		if val.v == v2 {
			return true
		}
	}
	return false
}

func (am *AdjacencyArray) Degree(v int) int {
	degree := len(am.m[v])
	for i := 0; i < degree; i++ {
		if am.m[v][i].v == v {
			return degree + 1
		}
	}

	return degree
}

func NewAdjacencyArray(countVertex int) *AdjacencyArray {
	m := make([][]*AdjacencyVertex, countVertex+1)
	return &AdjacencyArray{m: m}
}
