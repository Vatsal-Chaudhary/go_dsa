package graph

import (
	"testing"

	dsgraph "go_dsa/ds/graph"

	"github.com/stretchr/testify/assert"
)

func Test_Kosaraju(t *testing.T) {
	g1 := dsgraph.NewAdjacencyListGraph(dsgraph.DirectedUnweighted)

	g1.AddEdge(0, 2)
	g1.AddEdge(1, 0)
	g1.AddEdge(2, 1)
	g1.AddEdge(2, 3)
	g1.AddEdge(3, 4)
	g1.AddEdge(4, 3)

	res := normalizeComponents(KosarajuSCC(g1))

	assert.Equal(t, []int{0, 1, 2}, res[0])
	assert.Equal(t, []int{3, 4}, res[1])

	g2 := dsgraph.NewAdjacencyListGraph(dsgraph.DirectedUnweighted)

	g2.AddEdge(0, 1)
	g2.AddEdge(1, 2)
	g2.AddEdge(2, 1)
	g2.AddEdge(1, 3)
	g2.AddEdge(3, 4)
	g2.AddEdge(4, 5)
	g2.AddEdge(5, 4)

	res = normalizeComponents(KosarajuSCC(g2))

	assert.Equal(t, []int{0}, res[0])
	assert.Equal(t, []int{1, 2}, res[1])
	assert.Equal(t, []int{3}, res[2])
	assert.Equal(t, []int{4, 5}, res[3])

	g3 := dsgraph.NewAdjacencyListGraph(dsgraph.DirectedUnweighted)

	g3.AddEdge(0, 1)
	g3.AddEdge(1, 2)
	g3.AddEdge(2, 0)
	g3.AddEdge(3, 1)
	g3.AddEdge(3, 4)

	res = normalizeComponents(KosarajuSCC(g3))

	assert.Equal(t, []int{0, 1, 2}, res[0])
	assert.Equal(t, []int{3}, res[1])
	assert.Equal(t, []int{4}, res[2])
}
