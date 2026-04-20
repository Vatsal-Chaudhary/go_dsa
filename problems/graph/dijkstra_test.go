package graph

import (
	"math"
	"testing"

	dsgraph "go_dsa/ds/graph"

	"github.com/stretchr/testify/assert"
)

func Test_Dijkstra(t *testing.T) {
	g := dsgraph.NewAdjacencyListGraph(dsgraph.UndirectedWeighted)

	g.AddEdge(0, 1, 5)
	g.AddEdge(0, 2, 10)
	g.AddEdge(1, 2, 3)
	g.AddEdge(2, 3, 2)
	g.AddEdge(1, 3, 20)

	dist := Dijkstra_Shortest_Path(g, 0)

	assert.Equal(t, 0, dist[0])
	assert.Equal(t, 10, dist[3])
	assert.Equal(t, 8, dist[2])
}

func Test_Dijkstra_Unreachable(t *testing.T) {
	g := dsgraph.NewAdjacencyListGraph(dsgraph.UndirectedWeighted)

	g.AddEdge(0, 1, 5)
	g.AddEdge(2, 3, 10)

	dist := Dijkstra_Shortest_Path(g, 0)

	assert.Equal(t, 0, dist[0])
	assert.Equal(t, 5, dist[1])

	assert.Equal(t, math.MaxInt, dist[3])
	assert.Equal(t, math.MaxInt, dist[2])
}

func Test_Dijkstra_Directed(t *testing.T) {
	g := dsgraph.NewAdjacencyListGraph(dsgraph.DirectedWeighted)

	g.AddEdge(0, 1, 5)
	g.AddEdge(1, 2, 3)
	g.AddEdge(0, 2, 100)

	dist := Dijkstra_Shortest_Path(g, 0)

	assert.Equal(t, 8, dist[2])
}
