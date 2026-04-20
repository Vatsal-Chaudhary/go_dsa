package graph

import (
	"testing"

	dsgraph "go_dsa/ds/graph"

	"github.com/stretchr/testify/assert"
)

func TestBFS_FullyConnected(t *testing.T) {
	g := dsgraph.NewAdjacencyListGraph(dsgraph.UndirectedUnweighted)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	g.AddEdge(3, 5)

	result := BFS(g, 1)

	assert.Equal(t, 1, result[0])
	assert.Len(t, result, 5)

	pos := func(val int) int {
		for i, v := range result {
			if v == val {
				return i
			}
		}
		return -1
	}

	assert.Less(t, pos(2), pos(3))
	assert.Less(t, pos(2), pos(4))
	assert.Less(t, pos(3), pos(5))
}

func TestBFSDisconnected_TwoComponents(t *testing.T) {
	g := dsgraph.NewAdjacencyListGraph(dsgraph.UndirectedUnweighted)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(4, 5)

	result := normalizeComponents(BFSDisconnected(g))

	assert.Len(t, result, 2)
	assert.Equal(t, []int{1, 2, 3}, result[0])
	assert.Equal(t, []int{4, 5}, result[1])
}

func TestBFSDisconnected_ThreeComponents(t *testing.T) {
	g := dsgraph.NewAdjacencyListGraph(dsgraph.UndirectedUnweighted)
	g.AddEdge(1, 2)
	g.AddEdge(3, 4)
	g.AddEdge(5, 6)

	result := normalizeComponents(BFSDisconnected(g))

	assert.Len(t, result, 3)
	assert.Equal(t, []int{1, 2}, result[0])
	assert.Equal(t, []int{3, 4}, result[1])
	assert.Equal(t, []int{5, 6}, result[2])
}

func TestBFSCount(t *testing.T) {
	g := dsgraph.NewAdjacencyListGraph(dsgraph.UndirectedUnweighted)
	g.AddEdge(1, 2)
	g.AddEdge(3, 4)
	g.AddEdge(5, 6)

	count := BFSCount(g)

	assert.Equal(t, count, 3)

	g.AddEdge(5, 7)
	g.AddEdge(8, 9)
	count = BFSCount(g)
	assert.Equal(t, count, 4)
}
