package graph

import (
	"testing"
	dsgraph "go_dsa/ds/graph"
	"github.com/stretchr/testify/assert"
)

func Test_TarjanSCC(t *testing.T) {
	// same graph as kosaraju tests so you can compare results
	g1 := dsgraph.NewAdjacencyListGraph(dsgraph.DirectedUnweighted)
	g1.AddEdge(0, 2)
	g1.AddEdge(1, 0)
	g1.AddEdge(2, 1)
	g1.AddEdge(2, 3)
	g1.AddEdge(3, 4)
	g1.AddEdge(4, 3)

	res := normalizeComponents(TarjanSCC(g1))
	assert.Len(t, res, 2)
	assert.Equal(t, []int{0, 1, 2}, res[0])
	assert.Equal(t, []int{3, 4}, res[1])

	// linear chain — every vertex is its own SCC
	g2 := dsgraph.NewAdjacencyListGraph(dsgraph.DirectedUnweighted)
	g2.AddEdge(0, 1)
	g2.AddEdge(1, 2)
	g2.AddEdge(2, 3)

	res = normalizeComponents(TarjanSCC(g2))
	assert.Len(t, res, 4)

	// fully connected cycle — one SCC
	g3 := dsgraph.NewAdjacencyListGraph(dsgraph.DirectedUnweighted)
	g3.AddEdge(0, 1)
	g3.AddEdge(1, 2)
	g3.AddEdge(2, 0)

	res = normalizeComponents(TarjanSCC(g3))
	assert.Len(t, res, 1)
	assert.Equal(t, []int{0, 1, 2}, res[0])

	// same as kosaraju g2
	g4 := dsgraph.NewAdjacencyListGraph(dsgraph.DirectedUnweighted)
	g4.AddEdge(0, 1)
	g4.AddEdge(1, 2)
	g4.AddEdge(2, 1)
	g4.AddEdge(1, 3)
	g4.AddEdge(3, 4)
	g4.AddEdge(4, 5)
	g4.AddEdge(5, 4)

	res = normalizeComponents(TarjanSCC(g4))
	assert.Len(t, res, 4)
	assert.Equal(t, []int{0}, res[0])
	assert.Equal(t, []int{1, 2}, res[1])
	assert.Equal(t, []int{3}, res[2])
	assert.Equal(t, []int{4, 5}, res[3])
}
