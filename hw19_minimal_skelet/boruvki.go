package hw19

type Boruvki struct {
	Kraskal
	parents []int
}

func (b *Boruvki) initParents() {
	b.parents = make([]int, len(b.graph))
	for i := len(b.parents) - 1; i >= 0; i-- {
		b.parents[i] = i
	}
}

func (b *Boruvki) find(v int) int {
	if b.parents[v] == v {
		return v
	}
	b.parents[v] = b.find(b.parents[v])
	return b.parents[v]
}

func (b *Boruvki) merge(v1, v2 int) {
	rootV1 := b.find(v1)
	rootV2 := b.find(v2)
	b.parents[rootV2] = rootV1
}

func (b *Boruvki) GetSkeleton() ([]Edge, error) {
	if b.hasZeroVertex() {
		return nil, ErrMinimalSkeletonNotExists
	}

	sortEdges := b.getSortEdges()
	result := make([]Edge, 0, len(sortEdges))
	b.initParents()

	countEdges := 0
	countStop := len(b.startGraph) - 1
	for _, wEdge := range sortEdges {
		if b.find(wEdge.v1) == b.find(wEdge.v2) {
			continue
		}

		b.startGraph[wEdge.v1] = append(b.startGraph[wEdge.v1], wEdge.v2)
		b.startGraph[wEdge.v2] = append(b.startGraph[wEdge.v2], wEdge.v1)
		result = append(result, Edge{v1: wEdge.v1, v2: wEdge.v2})
		b.merge(wEdge.v1, wEdge.v2)
		countEdges++
		if countEdges >= countStop {
			break
		}
	}
	return result, nil
}

func NewBoruvki(graph [][]*GraphEdge) *Boruvki {
	return &Boruvki{
		Kraskal: Kraskal{
			graph:      graph,
			startGraph: make([][]int, len(graph)),
		},
	}
}
