package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUndirectedWeightedGraph(t *testing.T) {
	g := NewAdjacencyListGraph(UndirectedWeighted)

	assert.NotPanics(t, func() {
		g.AddEdge(1, 2, 4)
		g.AddEdge(1, 3, 5)
	})

	assert.Panics(t, func() {
		g.AddEdge(3, 5)
	})

	n1 := g.Neighbours(1)
	n2 := g.Neighbours(2)

	assert.Len(t, n1, 2)
	assert.Len(t, n2, 1)

	assert.ElementsMatch(t,
		[]int{2, 3},
		[]int{n1[0].To, n1[1].To},
	)

	assert.Equal(t, 4, n2[0].Weight)

	assert.ElementsMatch(t, []int{1, 2, 3}, g.Vertices())
}

func TestUndirectedUnweightedGraph(t *testing.T) {
	g := NewAdjacencyListGraph(UndirectedUnweighted)

	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 1)
	g.AddEdge(2, 4)

	n3 := g.Neighbours(3)
	n2 := g.Neighbours(2)

	assert.Len(t, n3, 2)
	assert.Len(t, n2, 3)

	for _, e := range n3 {
		assert.Equal(t, 1, e.Weight)
	}

	assert.ElementsMatch(t, []int{1, 2, 3, 4}, g.Vertices())
}

func TestDirectedUnweightedGraph(t *testing.T) {
	g := NewAdjacencyListGraph(DirectedUnweighted)

	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	g.AddEdge(3, 1)
	g.AddEdge(2, 4)

	n3 := g.Neighbours(3)
	n2 := g.Neighbours(2)

	assert.Len(t, n3, 1)
	assert.Len(t, n2, 2)

	for _, e := range n3 {
		assert.Equal(t, 1, e.Weight)
	}

	for _, e := range n2 {
		assert.NotEqual(t, 1, e.To)
	}

	assert.ElementsMatch(t, []int{1, 2, 3, 4}, g.Vertices())
}

func TestDirectedWeightedGraph(t *testing.T) {
	g := NewAdjacencyListGraph(DirectedWeighted)

	assert.NotPanics(t, func() {
		g.AddEdge(1, 2, 4)
		g.AddEdge(1, 3, 6)
		g.AddEdge(3, 4, 10)
	})

	assert.Panics(t, func() {
		g.AddEdge(4, 5)
	})

	n1 := g.Neighbours(1)
	n2 := g.Neighbours(2)
	n3 := g.Neighbours(3)

	assert.Len(t, n1, 2)
	assert.Len(t, n2, 0)
	assert.Len(t, n3, 1)

	assert.ElementsMatch(t,
		[]int{2, 3},
		[]int{n1[0].To, n1[1].To},
	)

	assert.Equal(t, n1[0].Weight, 4)

	assert.ElementsMatch(t, []int{1, 2, 3, 4}, g.Vertices())
}
