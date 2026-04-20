package graph

import (
	dsgraph "go_dsa/ds/graph"
	"math"
)

// Shortest path in undirected unweighted graph
func ShortestPath_UU(g dsgraph.Graph, src, dst int) int {
	vertices := g.Vertices()

	dist := make(map[int]int)
	for _, v := range vertices {
		dist[v] = math.MaxInt
	}

	dist[src] = 0

	visited := make(map[int]bool)
	visited[src] = true

	queue := []int{src}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		for _, edge := range g.Neighbours(node) {
			if !visited[edge.To] {
				dist[edge.To] = dist[node] + 1
				visited[edge.To] = true
				queue = append(queue, edge.To)
			}
		}
	}

	if dist[dst] == math.MaxInt {
		return -1
	}

	return dist[dst]
}

// LeetCode 1091 style — shortest path in unweighted graph
// returns -1 if no path exists
func ShortestPath_UU_Another(g dsgraph.Graph, src, dst int) int {
	if dst == src {
		return 0
	}

	visited := make(map[int]bool)
	visited[src] = true
	queue := []int{src}

	dist := map[int]int{src: 0}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		for _, edge := range g.Neighbours(node) {
			if !visited[edge.To] {
				visited[edge.To] = true
				dist[edge.To] = dist[node] + 1

				if edge.To == dst {
					return dist[edge.To]
				}

				queue = append(queue, edge.To)
			}
		}
	}

	return -1
}

// Shortest path in direct acyclic weighted graph
// O(V + E)
func ShortestPath_DAG(g dsgraph.Graph, src, dst int) int {
	vertices := g.Vertices()

	dist := make(map[int]int)
	for _, v := range vertices {
		dist[v] = math.MaxInt
	}

	dist[src] = 0

	top_sort := TopologicalSort_BFS(g)

	for _, u := range top_sort {
		for _, edge := range g.Neighbours(u) {
			if dist[edge.To] > dist[u]+edge.Weight {
				dist[edge.To] = dist[u] + edge.Weight
			}
		}
	}

	return dist[dst]
}
