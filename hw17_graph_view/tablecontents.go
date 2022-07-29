package graphview

type TableContents struct {
	vertex   []int
	contents []*AdjacencyVertex
}

func (am *TableContents) setAdjacencyVertex(edges []edge) {
	current := 0
	am.contents = append(am.contents, nil)

	for i := 1; i < len(am.vertex); i++ {
		added := 0
		for _, val := range edges {
			if i != val.v1 && i != val.v2 {
				continue
			}

			v := val.v2
			if i == val.v2 {
				v = val.v1
			}
			am.contents = append(am.contents, &AdjacencyVertex{
				v:      v,
				weight: val.weight,
			})
			added++
		}

		if added == 0 {
			am.vertex[i] = 0
		} else {
			am.vertex[i] = current + 1
			current += added
		}
	}
}

func (am *TableContents) getNextIndex(k int) int {
	k1 := len(am.vertex) - 1
	for i := k + 1; i < k1; i++ {
		if am.vertex[i] != 0 {
			return am.vertex[i]
		}
	}
	return len(am.contents)
}

func (am *TableContents) ListAdjacency(v int) []int {
	var result []int

	if am.vertex[v] == 0 {
		return result
	}

	end := am.getNextIndex(v)
	start := am.vertex[v]

	for i := start; i < end; i++ {
		result = append(result, am.contents[i].v)
	}
	return result
}

func (am *TableContents) IsAdjacency(v1, v2 int) bool {
	if am.vertex[v1] == 0 {
		return false
	}

	start := am.vertex[v1]

	k1 := am.getNextIndex(v1)
	for i := start; i < k1; i++ {
		if v2 == am.contents[i].v {
			return true
		}
	}

	return false
}

func (am *TableContents) Degree(v int) int {
	if am.vertex[v] == 0 {
		return 0
	}

	end := am.getNextIndex(v)

	degree := end - am.vertex[v]
	for i := am.vertex[v]; i < end; i++ {
		if v == am.contents[i].v {
			return degree + 1
		}
	}
	return degree
}

func NewTableContents(countVertex int) *TableContents {
	m := make([]int, countVertex+1)
	return &TableContents{vertex: m}
}
