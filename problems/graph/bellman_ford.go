package graph

import (
	dsgraph "go_dsa/ds/graph"
	"math"
)

func bellmanford_algo(g dsgraph.Graph, src int) map[int]int {
	dist := make(map[int]int)

	for _, v := range g.Vertices() {
		dist[v] = math.MaxInt
	}

	dist[src] = 0

	v := len(g.Vertices())
	for i := 0; i < v-1; i++ {
		for _, u := range g.Vertices() {
			for _, edge := range g.Neighbours(u) {
				if dist[u] != math.MaxInt && dist[edge.To] > dist[u]+edge.Weight {
					dist[edge.To] = dist[u] + edge.Weight
				}
			}
		}
	}

	return dist
}

func Bellmanford(g dsgraph.Graph, src int) map[int]int {
	return bellmanford_algo(g, src)
}

func DetectNegative(g dsgraph.Graph, src int) bool {
	dist := bellmanford_algo(g, src)

	for _, u := range g.Vertices() {
		for _, edge := range g.Neighbours(u) {
			if dist[u] != math.MaxInt && dist[edge.To] > dist[u]+edge.Weight {
				return true
			}
		}
	}
	return false
}
