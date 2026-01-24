package graph

type AdjacencyListGraph struct {
	adj      map[int][]Edge
	directed bool
	weighted bool
}

func NewAdjacencyListGraph(t GraphType) Graph {
	g := &AdjacencyListGraph{
		adj: make(map[int][]Edge),
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

func (g *AdjacencyListGraph) addVertex(v int) {
	if _, exists := g.adj[v]; !exists {
		g.adj[v] = []Edge{}
	}
}

func (g *AdjacencyListGraph) AddEdge(u, v int, w ...int) {
	weight := 1
	if g.weighted {
		if len(w) == 0 {
			panic("wighted graph requires weight")
		}
		weight = w[0]
	}
	g.addVertex(u)
	g.addVertex(v)

	g.adj[u] = append(g.adj[u], Edge{To: v, Weight: weight})

	if !g.directed {
		g.adj[v] = append(g.adj[v], Edge{To: u, Weight: weight})
	}
}

func (g *AdjacencyListGraph) Neighbours(v int) []Edge {
	return g.adj[v]
}

func (g *AdjacencyListGraph) Vertices() []int {
	vertices := make([]int, 0, len(g.adj))
	for v := range g.adj {
		vertices = append(vertices, v)
	}

	return vertices
}
