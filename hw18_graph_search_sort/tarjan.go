package hw18

const (
	white = 0
	gray  = 1
	black = 2
)

type Tarjan struct {
	graph   [][]int
	paints  []int
	result  []int
	current int
}

func (t *Tarjan) Sort() ([]int, error) {
	t.result = make([]int, len(t.graph))
	t.current = len(t.graph) - 1

	for i := len(t.graph) - 1; i >= 0 && t.current >= 0; i-- {
		if t.paints[i] != white {
			continue
		}
		dfsResult := t.dfs(i)
		if !dfsResult {
			return nil, ErrSortNotExists
		}
	}

	return t.result, nil
}

func (t *Tarjan) dfs(v int) bool {
	t.paints[v] = gray
	for _, v1 := range t.graph[v] {
		if v1 == -1 {
			break
		}
		if t.paints[v1] == gray {
			return false
		}
		if t.paints[v1] == black {
			continue
		}
		if !t.dfs(v1) {
			return false
		}
	}
	t.paints[v] = black
	t.result[t.current] = v
	t.current--

	return true
}

func NewTarjan(graph [][]int) *Tarjan {
	return &Tarjan{
		graph:  graph,
		paints: make([]int, len(graph)),
	}
}
