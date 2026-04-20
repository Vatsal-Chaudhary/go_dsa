package graph

import (
	dsgraph "go_dsa/ds/graph"
	"math"
)

// Minimize wire length and make sure that all the computers are connected to each other
// may be through intermediate computers.
// Using Prims Algo

func PrimMST(g dsgraph.Graph) int {
	vertices := g.Vertices()
	v := len(vertices)

	key := make(map[int]int)
	inMST := make(map[int]bool)

	for i, _ := range vertices {
		key[i] = math.MaxInt
	}

	key[vertices[0]] = 0
	count := 0
	res := 0

	for count < v {
		u := -1

		for _, vertex := range vertices {
			if !inMST[vertex] && (u == -1 || key[vertex] < key[u]) {
				u = vertex
			}
		}

		inMST[u] = true
		res += key[u]

		for _, edge := range g.Neighbours(u) {
			if !inMST[edge.To] && edge.Weight < key[edge.To] {
				key[edge.To] = edge.Weight
			}
		}

		count++
	}

	return res
}
