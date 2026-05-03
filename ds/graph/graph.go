package graph

type GraphType int

const (
	UndirectedUnweighted GraphType = iota
	UndirectedWeighted
	DirectedUnweighted
	DirectedWeighted
)

type Edge struct {
	To     int
	Weight int
}

type Graph interface {
	AddVertex(v int)
	AddEdge(u, v int, w ...int)
	Neighbours(v int) []Edge
	Vertices() []int
}
