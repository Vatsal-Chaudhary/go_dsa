package graph

type AdjacencyMatrixGraph struct {
	matrix   [][]int
	directed bool
	weighted bool
	n        int
}

func NewAdjacencyMatrixGraph(n int, t GraphType) Graph {
	m := make([][]int, n)
	for i := 0; i < n; i++ {
		m[i] = make([]int, n)
	}

	g := &AdjacencyMatrixGraph{
		matrix: m,
		n:      n,
	}

	switch t {
	case UndirectedUnweighted:
		g.directed = false
		g.weighted = false

	case DirectedWeighted:
		g.directed = true
		g.weighted = true

	case DirectedUnweighted:
		g.directed = true
		g.weighted = false

	case UndirectedWeighted:
		g.directed = false
		g.weighted = true
	}

	return g
}

func (g *AdjacencyMatrixGraph) AddEdge(u, v int, w ...int) {
	if u < 0 || v < 0 || u >= g.n || v >= g.n {
		panic("vertex index out of bounds")
	}

	weight := 1
	if g.weighted {
		if len(w) == 0 {
			panic("weighted graph require weight")
		}
		weight = w[0]
	}

	g.matrix[u][v] = weight

	if !g.directed {
		g.matrix[v][u] = weight
	}
}

func (g *AdjacencyMatrixGraph) Neighbours(u int) []Edge {
	if u < 0 || u >= g.n {
		panic("vertix index out of bound")
	}
	var edge []Edge
	for v := 0; v < g.n; v++ {
		if g.matrix[u][v] != 0 {
			edge = append(edge, Edge{
				To:     v,
				Weight: g.matrix[u][v],
			})
		}
	}

	return edge
}

func (g *AdjacencyMatrixGraph) Vertices() []int {
	vs := make([]int, g.n)
	for i := 0; i < g.n; i++ {
		vs[i] = i
	}
	return vs
}

func (g *AdjacencyMatrixGraph) AddVertex(v int) {

}
