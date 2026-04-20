package graph

import (
	"testing"

	dsgraph "go_dsa/ds/graph"

	"github.com/stretchr/testify/assert"
)

func TestDetectsCycle_Undirected(t *testing.T) {
	graph := dsgraph.NewAdjacencyListGraph(dsgraph.UndirectedUnweighted)

	graph.AddEdge(1, 2)
	graph.AddEdge(2, 3)
	graph.AddEdge(3, 4)

	isCycle := DetectsCycle_Undirected(graph)

	assert.False(t, isCycle)

	graph.AddEdge(4, 1)
	isCycle = DetectsCycle_Undirected(graph)

	assert.True(t, isCycle)
}

func TestDetectsCycle_Directed(t *testing.T) {
	g1 := dsgraph.NewAdjacencyListGraph(dsgraph.DirectedUnweighted)

	g1.AddEdge(1, 2)
	g1.AddEdge(2, 3)
	g1.AddEdge(3, 4)

	isCycle := DetectsCycle_Directed(g1)

	assert.False(t, isCycle)

	g1.AddEdge(4, 1)
	isCycle = DetectsCycle_Directed(g1)

	assert.True(t, isCycle)

	g2 := dsgraph.NewAdjacencyListGraph(dsgraph.DirectedUnweighted)

	g2.AddEdge(0, 1)
	g2.AddEdge(1, 3)
	g2.AddEdge(2, 3)
	g2.AddEdge(2, 1)

	assert.False(t, DetectsCycle_Directed(g2))

	g2.AddEdge(1, 4)
	g2.AddEdge(4, 0)
	assert.True(t, DetectsCycle_Directed(g2))
}

func TestDetectCycle_Directed_UsingKahns(t *testing.T) {
	g1 := dsgraph.NewAdjacencyListGraph(dsgraph.DirectedUnweighted)

	g1.AddEdge(1, 2)
	g1.AddEdge(2, 3)
	g1.AddEdge(3, 4)

	isCycle := DetectsCycle_Directed(g1)

	assert.False(t, isCycle)

	g1.AddEdge(4, 1)
	isCycle = DetectsCycle_Directed(g1)

	assert.True(t, isCycle)

	g2 := dsgraph.NewAdjacencyListGraph(dsgraph.DirectedUnweighted)

	g2.AddEdge(0, 1)
	g2.AddEdge(1, 3)
	g2.AddEdge(2, 3)
	g2.AddEdge(2, 1)

	assert.False(t, DetectsCycle_Directed(g2))

	g2.AddEdge(1, 4)
	g2.AddEdge(4, 0)
	assert.True(t, DetectsCycle_Directed(g2))
}
