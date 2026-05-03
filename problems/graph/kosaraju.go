package graph

import (
	dsgraph "go_dsa/ds/graph"
)

// Kosaraju's Algorithm
// Used to find out Strongly Connected Components
// Strongly Connected Components (SCC) -- Subgraph where every vertex is reachable from every other vertex

func KosarajuSCC(g dsgraph.Graph) [][]int {
	// step 1
	visited := make(map[int]bool)

	stack := []int{}

	for _, v := range g.Vertices() {
		if !visited[v] {
			dfs_from(g, &stack, visited, v)
		}
	}

	// step 2
	reveresd := dsgraph.NewAdjacencyListGraph(dsgraph.DirectedUnweighted)

	for _, u := range g.Vertices() {
		for _, edge := range g.Neighbours(u) {
			reveresd.AddEdge(edge.To, u)
		}
	}

	// step 3
	visited1 := make(map[int]bool)

	result := [][]int{}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		res := []int{}

		if !visited1[node] {
			dfs_from(reveresd, &res, visited1, node)
			result = append(result, res)
		}

	}

	return result
}

func dfs_from(g dsgraph.Graph, stack *[]int, visited map[int]bool, v int) {
	visited[v] = true

	for _, edge := range g.Neighbours(v) {
		if !visited[edge.To] {
			dfs_from(g, stack, visited, edge.To)
		}
	}

	*stack = append(*stack, v)
}
