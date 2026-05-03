package graph

import (
	"math"

	dsgraph "go_dsa/ds/graph"
)

// Time -- O(v * v)
func Dijkstra_Shortest_Path(g dsgraph.Graph, src int) map[int]int {
	vertices := g.Vertices()
	v := len(vertices)

	dist := make(map[int]int)

	for _, vertex := range vertices {
		dist[vertex] = math.MaxInt
	}

	dist[src] = 0

	visited := make(map[int]bool)

	count := 0

	for count < v {
		u := -1

		for _, vertex := range vertices {
			if !visited[vertex] && (u == -1 || dist[u] > dist[vertex]) {
				u = vertex
			}
		}

		visited[u] = true

		if dist[u] == math.MaxInt {
			break
		}

		for _, edge := range g.Neighbours(u) {
			if !visited[edge.To] && dist[edge.To] > dist[u]+edge.Weight {
				dist[edge.To] = dist[u] + edge.Weight
			}
		}

		count++
	}

	return dist
}
