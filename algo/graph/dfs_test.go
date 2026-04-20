package graph

import (
	"testing"

	dsgraph "go_dsa/ds/graph"

	"github.com/stretchr/testify/assert"
)

func TestDFS_SimpleUndirected(t *testing.T) {
	g := dsgraph.NewAdjacencyListGraph(dsgraph.UndirectedUnweighted)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)
	g.AddEdge(3, 5)

	result := DFS(g, 1)

	assert.Len(t, result, 5)
	assert.Equal(t, result[0], 1)

	pos := func(val int) int {
		for i, v := range result {
			if v == val {
				return i
			}
		}
		return -1
	}

	assert.Less(t, pos(3), pos(5))
	assert.Less(t, pos(2), pos(4))
	assert.Less(t, pos(4), pos(3))
	assert.Less(t, pos(4), pos(5))
}

func TestDFS_DisconnectedGraph(t *testing.T) {
	graph := dsgraph.NewAdjacencyListGraph(dsgraph.UndirectedUnweighted)

	graph.AddEdge(1, 2)
	graph.AddEdge(3, 4)

	result := DFS(graph, 1)

	assert.Len(t, result, 2)
	assert.Contains(t, result, 1)
	assert.Contains(t, result, 2)
	assert.NotContains(t, result, 3)
	assert.NotContains(t, result, 4)
}

func TestDFSDisconnected_TwoComponents(t *testing.T) {
	graph := dsgraph.NewAdjacencyListGraph(dsgraph.UndirectedUnweighted)

	graph.AddEdge(1, 2)
	graph.AddEdge(2, 3)
	graph.AddEdge(4, 5)

	result := normalizeComponents(DFSDisconnected(graph))

	assert.Len(t, result, 2)
	assert.Equal(t, []int{1, 2, 3}, result[0])
	assert.Equal(t, []int{4, 5}, result[1])
}

func TestDFSDisconnected_ThreeComponents(t *testing.T) {
	graph := dsgraph.NewAdjacencyListGraph(dsgraph.UndirectedUnweighted)

	graph.AddEdge(1, 2)
	graph.AddEdge(3, 4)
	graph.AddEdge(5, 6)

	result := normalizeComponents(DFSDisconnected(graph))

	assert.Len(t, result, 3)
	assert.Equal(t, []int{1, 2}, result[0])
	assert.Equal(t, []int{3, 4}, result[1])
	assert.Equal(t, []int{5, 6}, result[2])
}

func TestDFS_FullyConnected(t *testing.T) {
	graph := dsgraph.NewAdjacencyListGraph(dsgraph.UndirectedUnweighted)

	graph.AddEdge(1, 2)
	graph.AddEdge(2, 3)
	graph.AddEdge(3, 4)

	result := normalizeComponents(DFSDisconnected(graph))

	assert.Len(t, result, 1)
	assert.Len(t, result[0], 4)
}

func TestDFS_Count(t *testing.T) {
	graph := dsgraph.NewAdjacencyListGraph(dsgraph.UndirectedUnweighted)

	graph.AddEdge(1, 2)
	graph.AddEdge(2, 3)
	graph.AddEdge(4, 5)
	graph.AddEdge(6, 7)

	count := DFSCount(graph)

	assert.Equal(t, count, 3)

	graph.AddEdge(8, 9)
	count = DFSCount(graph)
	assert.Equal(t, count, 4)
}
