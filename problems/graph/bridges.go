package graph

import (
	dsgraph "go_dsa/ds/graph"
)

// Finding bridges in a graph using tarjans algorithm
func Bridges(g dsgraph.Graph) [][2]int {
	visited := make(map[int]bool)
	disc := make(map[int]int)
	low := make(map[int]int)
	parent := make(map[int]int)
	timer := 0
	result := [][2]int{}

	for _, v := range g.Vertices() {
		if !visited[v] {
			parent[v] = -1
			dfs_bridge(g, v, visited, disc, low, parent, &timer, &result)
		}
	}

	return result
}

func dfs_bridge(
	g dsgraph.Graph,
	u int,
	visited map[int]bool,
	disc map[int]int,
	low map[int]int,
	parent map[int]int,
	timer *int,
	result *[][2]int,
) {
	visited[u] = true

	*timer++
	low[u] = *timer
	disc[u] = *timer

	for _, edge := range g.Neighbours(u) {
		v := edge.To
		if !visited[edge.To] {

			parent[v] = u

			dfs_bridge(g, v, visited, disc, low, parent, timer, result)

			low[u] = min(low[u], low[v])

			if disc[u] < low[v] {
				u1, v1 := min(u, v), max(u, v)
				*result = append(*result, [2]int{u1, v1})
			}

		} else if v != parent[u] {
			low[u] = min(low[u], disc[v])
		}
	}
}
