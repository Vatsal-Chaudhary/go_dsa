package graph

import (
	dsgraph "go_dsa/ds/graph"
)

// Undirected helper
func dfs_rec_undirected(g dsgraph.Graph, visited map[int]bool, s, parent int) bool {
	visited[s] = true

	for _, edge := range g.Neighbours(s) {
		if !visited[edge.To] {
			if dfs_rec_undirected(g, visited, edge.To, s) == true {
				return true
			}
		} else if edge.To != parent {
			return true
		}
	}

	return false
}

// Directed Helper
func dfs_rec_directed(g dsgraph.Graph, visited map[int]bool, recSt map[int]bool, src int) bool {
	visited[src] = true
	recSt[src] = true

	for _, edge := range g.Neighbours(src) {
		if !visited[edge.To] {
			if dfs_rec_directed(g, visited, recSt, edge.To) {
				return true
			}
		} else if recSt[edge.To] == true {
			return true
		}
	}

	recSt[src] = false

	return false
}

// Detects cycle in undirected graph
func DetectsCycle_Undirected(g dsgraph.Graph) bool {
	visited := make(map[int]bool)

	for _, v := range g.Vertices() {
		if !visited[v] {
			if dfs_rec_undirected(g, visited, v, -1) == true {
				return true
			}
		}
	}
	return false
}

// Detects cycle in an directed graph
func DetectsCycle_Directed(g dsgraph.Graph) bool {
	visited := make(map[int]bool)
	recSt := make(map[int]bool)

	for _, v := range g.Vertices() {
		if !visited[v] {
			if dfs_rec_directed(g, visited, recSt, v) {
				return true
			}
		}
	}

	return false
}

// detects cycle detection direct graph using kahn's algo
func DetectsCycle_Directed_Kahn(g dsgraph.Graph) bool {
	inDegree := make([]int, len(g.Vertices()))

	for _, v := range g.Vertices() {
		for _, edge := range g.Neighbours(v) {
			inDegree[edge.To]++
		}
	}
	queue := []int{}

	for _, v := range g.Vertices() {
		if inDegree[v] == 0 {
			queue = append(queue, v)
		}
	}

	count := 0

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		for _, edge := range g.Neighbours(node) {
			inDegree[edge.To]--

			if inDegree[edge.To] == 0 {
				queue = append(queue, edge.To)
			}
		}

		count++
	}

	return count != len(g.Vertices())
}
