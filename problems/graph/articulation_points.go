package graph

import (
	"fmt"
	algo "go_dsa/algo/graph"
	dsgraph "go_dsa/ds/graph"
)

// Naive - O(V * (V+E))
func Articulation_points_naive(g dsgraph.Graph) []int {
	vertices := g.Vertices()
	original := len(algo.BFSDisconnected(g))
	result := []int{}
	fmt.Println("original", original)

	for _, skip := range vertices {

		temp := dsgraph.NewAdjacencyListGraph(dsgraph.UndirectedUnweighted)
		for _, u := range vertices {
			if u == skip {
				continue
			}
			temp.AddVertex(u)
		}

		added := make(map[[2]int]bool)
		for _, u := range vertices {
			if u == skip {
				continue
			}
			for _, edge := range g.Neighbours(u) {
				if edge.To != skip {
					key := [2]int{min(u, edge.To), max(u, edge.To)}
					if !added[key] {
						added[key] = true
						temp.AddEdge(u, edge.To)
					}
				}
			}
		}

		if len(algo.BFSDisconnected(temp)) > original {
			result = append(result, skip)
		}
	}

	return result
}

func ArticulationPoints(g dsgraph.Graph) []int {
	visited := make(map[int]bool)
	disc := make(map[int]int)
	low := make(map[int]int)
	parent := make(map[int]int)
	childCount := make(map[int]int)
	timer := 0
	result := []int{}
	ap := make(map[int]bool)

	for _, v := range g.Vertices() {
		if !visited[v] {
			parent[v] = -1
			dfs_ap(g, v, visited, disc, low, parent, &timer, ap, childCount)
		}
	}

	for v, isAP := range ap {
		if isAP {
			result = append(result, v)
		}
	}

	return result
}

func dfs_ap(
	g dsgraph.Graph,
	u int,
	visited map[int]bool,
	disc map[int]int,
	low map[int]int,
	parent map[int]int,
	timer *int,
	ap map[int]bool,
	childCount map[int]int,
) {
	visited[u] = true

	*timer++
	low[u] = *timer
	disc[u] = *timer

	for _, edge := range g.Neighbours(u) {
		v := edge.To
		if !visited[edge.To] {

			parent[v] = u
			childCount[u]++

			dfs_ap(g, v, visited, disc, low, parent, timer, ap, childCount)

			low[u] = min(low[u], low[v])

			if parent[u] == -1 && childCount[u] >= 2 {
				ap[u] = true
			}

			if parent[u] != -1 && low[v] >= disc[u] {
				ap[u] = true
			}

		} else if v != parent[u] {
			low[u] = min(low[u], disc[v])
		}
	}
}
