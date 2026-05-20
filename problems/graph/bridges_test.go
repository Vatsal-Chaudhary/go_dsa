package graph

import (
	"testing"
	dsgraph "go_dsa/ds/graph"
	"github.com/stretchr/testify/assert"
)

func Test_Bridges(t *testing.T) {
	// simple chain — every edge is a bridge
	g1 := dsgraph.NewAdjacencyListGraph(dsgraph.UndirectedUnweighted)
	g1.AddEdge(0, 1)
	g1.AddEdge(1, 2)
	g1.AddEdge(2, 3)
	res := Bridges(g1)
	assert.Len(t, res, 3)

	// triangle — no bridges (cycle covers every edge)
	g2 := dsgraph.NewAdjacencyListGraph(dsgraph.UndirectedUnweighted)
	g2.AddEdge(0, 1)
	g2.AddEdge(1, 2)
	g2.AddEdge(2, 0)
	res = Bridges(g2)
	assert.Empty(t, res)

	// two triangles connected by one edge
	g3 := dsgraph.NewAdjacencyListGraph(dsgraph.UndirectedUnweighted)
	g3.AddEdge(0, 1)
	g3.AddEdge(1, 2)
	g3.AddEdge(2, 0)
	g3.AddEdge(2, 3)  // ← this is the bridge
	g3.AddEdge(3, 4)
	g3.AddEdge(4, 5)
	g3.AddEdge(5, 3)
	res = Bridges(g3)
	assert.Len(t, res, 1)
	assert.Contains(t, res, [2]int{2, 3})

	// graph from AP problem — 0-1-2 triangle, then chain 2-3-4
	g4 := dsgraph.NewAdjacencyListGraph(dsgraph.UndirectedUnweighted)
	g4.AddEdge(0, 1)
	g4.AddEdge(0, 2)
	g4.AddEdge(1, 2)
	g4.AddEdge(2, 3)
	g4.AddEdge(3, 4)
	res = Bridges(g4)
	assert.Len(t, res, 2)
	assert.Contains(t, res, [2]int{2, 3})
	assert.Contains(t, res, [2]int{3, 4})
}


