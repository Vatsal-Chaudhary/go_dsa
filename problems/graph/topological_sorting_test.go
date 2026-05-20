package graph

import (
	"testing"

	dsgraph "go_dsa/ds/graph"

	"github.com/stretchr/testify/assert"
)

func TestTopologicalSort_BFS(t *testing.T) {
	g1 := dsgraph.NewAdjacencyListGraph(dsgraph.DirectedUnweighted)

	g1.AddEdge(0, 1)
	g1.AddEdge(0, 2)

	res1 := TopologicalSort_BFS(g1)

	assert.Equal(t, []int{0, 1, 2}, res1)

	g2 := dsgraph.NewAdjacencyListGraph(dsgraph.DirectedUnweighted)

	g2.AddEdge(0, 2)
	g2.AddEdge(0, 3)
	g2.AddEdge(1, 3)
	g2.AddEdge(1, 4)

	res2 := TopologicalSort_BFS(g2)

	pos := func(v int, res []int) int {
		for i, _ := range res {
			if res[i] == v {
				return i
			}
		}
		return -1
	}

	assert.Less(t, pos(0, res2), pos(3, res2))
	assert.Less(t, pos(1, res2), pos(3, res2))
}

func TestTopologicalSort_DFS(t *testing.T) {
	g1 := dsgraph.NewAdjacencyListGraph(dsgraph.DirectedUnweighted)

	g1.AddEdge(0, 1)
	g1.AddEdge(0, 2)

	res1 := TopologicalSort_BFS(g1)

	assert.Equal(t, []int{0, 1, 2}, res1)

	g2 := dsgraph.NewAdjacencyListGraph(dsgraph.DirectedUnweighted)

	g2.AddEdge(0, 2)
	g2.AddEdge(0, 3)
	g2.AddEdge(1, 3)
	g2.AddEdge(1, 4)

	res2 := TopologicalSort_DFS(g2)

	pos := func(v int, res []int) int {
		for i, _ := range res {
			if res[i] == v {
				return i
			}
		}
		return -1
	}

	assert.Less(t, pos(0, res2), pos(3, res2))
	assert.Less(t, pos(1, res2), pos(3, res2))
}
