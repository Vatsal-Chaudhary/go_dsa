package graph

import (
	"testing"

	dsgraph "go_dsa/ds/graph"

	"github.com/stretchr/testify/assert"
)

func Test_PrimMST(t *testing.T) {
	g := dsgraph.NewAdjacencyListGraph(dsgraph.UndirectedWeighted)

	g.AddEdge(0, 1, 2)
	g.AddEdge(0, 3, 6)
	g.AddEdge(1, 2, 3)
	g.AddEdge(1, 4, 5)
	g.AddEdge(2, 4, 7)
	g.AddEdge(3, 4, 9)

	assert.Equal(t, 16, PrimMST(g))

	g.AddEdge(2, 5, 6)
	g.AddEdge(4, 5, 5)

	assert.Equal(t, 21, PrimMST(g))
}
