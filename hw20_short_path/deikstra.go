package hw20

const notSet = -1

type Edge struct {
	v1 int
	v2 int
}

type GraphEdge struct {
	v1     int
	weight int
}

// Dijkstra алгоритм вычисления кратчайшего пути в графе с неотрицательными ребрами.
type Dijkstra struct {
	graph    [][]*GraphEdge
	distance []int
	prev     []int
	queue    []int
}

func (d *Dijkstra) init(vertex int) {
	for i := range d.distance {
		d.distance[i] = notSet
	}
	d.distance[vertex] = 0
	for i := range d.prev {
		d.prev[i] = notSet
	}
	d.queue = make([]int, 0, len(d.distance))
	d.queue = append(d.queue, vertex)
}

// dequeue возвращает индекс следующей обрабатываемой вершины  или - 1 если в очереди c приоритетами нет элементов.
func (d *Dijkstra) dequeue() int {
	min := notSet
	for i, v := range d.queue {
		if v == notSet {
			continue
		}
		if min == notSet {
			min = i
			continue
		}
		if d.distance[d.queue[min]] > d.distance[v] {
			min = i
		}
	}

	if min == notSet {
		return notSet
	}

	vertex := d.queue[min]
	d.queue[min] = notSet
	return vertex
}

func (d *Dijkstra) enqueue(v int) {
	d.queue = append(d.queue, v)
}

// GetMinPath возвращает минимальный путь (набор ребер и значение кратчайшего пути) из v1->v2.
func (d *Dijkstra) GetMinPath(v1, v2 int) ([]Edge, int) {
	if v1 == v2 {
		return nil, 0
	}
	d.init(v1)
	for {
		v := d.dequeue()
		if v == notSet || v == v2 {
			break
		}
		for _, graphEdge := range d.graph[v] {
			if graphEdge == nil {
				break
			}
			value := d.distance[v] + graphEdge.weight
			if d.distance[graphEdge.v1] != notSet && d.distance[graphEdge.v1] <= value {
				continue
			}

			if d.distance[graphEdge.v1] == notSet {
				d.enqueue(graphEdge.v1)
			}

			d.distance[graphEdge.v1] = value
			d.prev[graphEdge.v1] = v
		}
	}

	var result []Edge
	if d.distance[v2] == notSet {
		return nil, notSet
	}

	current := v2
	for current != v1 {
		result = append(result, Edge{
			v1: d.prev[current],
			v2: current,
		})
		current = d.prev[current]
	}

	lenResult := len(result)
	for i := lenResult/2 - 1; i >= 0; i-- {
		result[i], result[lenResult-i-1] = result[lenResult-i-1], result[i]
	}

	return result, d.distance[v2]
}

func NewDeikstra(graph [][]*GraphEdge) *Dijkstra {
	return &Dijkstra{
		graph:    graph,
		distance: make([]int, len(graph)),
		prev:     make([]int, len(graph)),
	}
}
