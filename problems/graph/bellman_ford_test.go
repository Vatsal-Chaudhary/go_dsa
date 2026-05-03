package graph

import (
	"testing"

	dsgraph "go_dsa/ds/graph"

	"github.com/stretchr/testify/assert"
)

func Test_BellmanFord(t *testing.T) {
	// simple path with positive weights
	g := dsgraph.NewAdjacencyListGraph(dsgraph.DirectedWeighted)

	g.AddEdge(0, 1, 4)
	g.AddEdge(0, 2, 2)
	g.AddEdge(1, 2, 1)
	g.AddEdge(2, 3, 5)
	g.AddEdge(1, 3, 2)

	res := Bellmanford(g, 0)

	assert.Equal(t, 0, res[0])
	assert.Equal(t, 4, res[1])
	assert.Equal(t, 2, res[2])
	assert.Equal(t, 6, res[3])

	// negative edges but no negative cycle
	g1 := dsgraph.NewAdjacencyListGraph(dsgraph.DirectedWeighted)
	g1.AddEdge(0, 1, 1)
	g1.AddEdge(0, 2, 4)
	g1.AddEdge(1, 2, -2)
	g1.AddEdge(2, 3, 2)
	res = Bellmanford(g1, 0)

	assert.Equal(t, 0, res[0])
	assert.Equal(t, 1, res[1])
	assert.Equal(t, -1, res[2])
	assert.Equal(t, 1, res[3])

}

func Test_BellmanFord_DetectCycle(t *testing.T) {
	// negative cycle detection
	g2 := dsgraph.NewAdjacencyListGraph(dsgraph.DirectedWeighted)
	g2.AddEdge(0, 1, 1)
	g2.AddEdge(1, 2, 3)
	g2.AddEdge(2, 1, -5)
	g2.AddEdge(2, 3, 2)

	isN := DetectNegative(g2, 0)

	assert.True(t, isN)
}
