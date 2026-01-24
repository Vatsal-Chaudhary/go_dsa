package main

import (
	"fmt"
	"go_dsa/ds/graph"
)

func main() {
	demoGraph()
}

func demoGraph() {
	// undirectedWeighted()
	undirectedUnweighted()
}

func undirectedWeighted() {
	g := graph.NewAdjacencyListGraph(graph.UndirectedWeighted)

	g.AddEdge(1, 2, 4)
	g.AddEdge(1, 3, 5)
	g.AddEdge(2, 3, 5)
	g.AddEdge(2, 4, 10)

	for _, v := range g.Vertices() {
		fmt.Printf("Vertex %d: ", v)
		for _, e := range g.Neighbours(v) {
			fmt.Printf("(%d, w=%d)", e.To, e.Weight)
		}
		fmt.Println()
	}
}

func undirectedUnweighted() {
	g := graph.NewAdjacencyListGraph(graph.UndirectedUnweighted)

	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)

	for _, v := range g.Vertices() {
		fmt.Printf("Vertex %d: ", v)
		for _, e := range g.Neighbours(v) {
			fmt.Printf("(%d, w=%d)", e.To, e.Weight)
		}
		fmt.Println()
	}
}
