package graph

import (
	"fmt"
	dsgraph "go_dsa/ds/graph"
)

func dfs_rec(graph dsgraph.Graph, visited map[int]bool, src int, result *[]int) []int {
	visited[src] = true

	fmt.Printf("%d \n", src)
	*result = append(*result, src)

	for _, edge := range graph.Neighbours(src) {
		if !visited[edge.To] {
			dfs_rec(graph, visited, edge.To, result)
		}
	}

	return *result
}

// From single source and Connected
func DFS(g dsgraph.Graph, src int) []int {
	visited := make(map[int]bool)
	result := []int{}

	dfs_rec(g, visited, src, &result)

	return result
}

// Disconnected all over graph
func DFSDisconnected(g dsgraph.Graph) [][]int {
	visited := make(map[int]bool)
	result := [][]int{}

	for _, v := range g.Vertices() {
		if !visited[v] {
			component := []int{}
			dfs_rec(g, visited, v, &component)
			result = append(result, component)
		}
	}

	return result
}

// Count connnected components
func DFSCount(g dsgraph.Graph) int {
	visited := make(map[int]bool)
	result := [][]int{}

	for _, v := range g.Vertices() {
		if !visited[v] {
			component := []int{}
			dfs_rec(g, visited, v, &component)
			result = append(result, component)
		}
	}

	return len(result)
}
