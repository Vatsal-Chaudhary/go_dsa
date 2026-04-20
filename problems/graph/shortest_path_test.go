package graph

import (
	dsgraph "go_dsa/ds/graph"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShortestPath_UU(t *testing.T) {
	graph := dsgraph.NewAdjacencyListGraph(dsgraph.UndirectedUnweighted)

	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(0, 4)
	graph.AddEdge(1, 3)
	graph.AddEdge(2, 3)
	graph.AddEdge(2, 4)
	graph.AddEdge(3, 5)
	graph.AddEdge(4, 5)

	dist := ShortestPath_UU(graph, 0, 4)

	assert.Equal(t, 1, dist)

	dist = ShortestPath_UU(graph, 2, 5)
	assert.Equal(t, 2, dist)
}

func Test_ShortestPath_DAG(t *testing.T) {
	graph := dsgraph.NewAdjacencyListGraph(dsgraph.DirectedWeighted)

	graph.AddEdge(0, 1, 2)
	graph.AddEdge(0, 4, 1)
	graph.AddEdge(1, 2, 3)
	graph.AddEdge(2, 3, 6)
	graph.AddEdge(4, 2, 2)
	graph.AddEdge(4, 5, 4)
	graph.AddEdge(5, 3, 1)

	dist := ShortestPath_DAG(graph, 0, 2)

	assert.Equal(t, 3, dist)

	dist = ShortestPath_DAG(graph, 0, 3)
	assert.Equal(t, 6, dist)
}
