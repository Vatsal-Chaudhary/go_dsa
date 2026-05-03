package graph

import (
	"testing"

	dsgraph "go_dsa/ds/graph"

	"github.com/stretchr/testify/assert"
)

func Test_ArticulationPoints(t *testing.T) {
	g1 := dsgraph.NewAdjacencyListGraph(dsgraph.UndirectedUnweighted)

	g1.AddEdge(0, 1)
	g1.AddEdge(0, 2)
	g1.AddEdge(1, 2)
	g1.AddEdge(2, 3)
	g1.AddEdge(3, 4)

	res := Articulation_points_naive(g1)

	assert.ElementsMatch(t, []int{2, 3}, res)

	g2 := dsgraph.NewAdjacencyListGraph(dsgraph.UndirectedUnweighted)

	g2.AddEdge(0, 1)
	g2.AddEdge(1, 2)
	g2.AddEdge(0, 2)
	g2.AddEdge(2, 3)
	g2.AddEdge(4, 3)
	g2.AddEdge(4, 2)

	res = Articulation_points_naive(g2)

	assert.ElementsMatch(t, []int{2}, res)

	g3 := dsgraph.NewAdjacencyListGraph(dsgraph.UndirectedUnweighted)

	g3.AddEdge(0, 2)
	g3.AddEdge(0, 1)
	g3.AddEdge(2, 1)

	res = Articulation_points_naive(g3)

	assert.Empty(t, res)
}
