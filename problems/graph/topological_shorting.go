package graph

import (
	dsgraph "go_dsa/ds/graph"
)

// Using Khans's algorithm (BFS based)
func TopologicalSort_BFS(g dsgraph.Graph) []int {
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

	result := []int{}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		result = append(result, node)

		for _, edge := range g.Neighbours(node) {
			inDegree[edge.To]--

			if inDegree[edge.To] == 0 {
				queue = append(queue, edge.To)
			}
		}
	}

	return result
}

// using DFS
func TopologicalSort_DFS(g dsgraph.Graph) []int {
	stack := []int{}
	visited := make(map[int]bool)

	for _, v := range g.Vertices() {
		if !visited[v] {
			dfs_top_sc(g, visited, &stack, v)
		}
	}

	result := []int{}
	for len(stack) > 0 {
		result = append(result, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return result
}

func dfs_top_sc(g dsgraph.Graph, visited map[int]bool, stack *[]int, u int) {
	visited[u] = true

	for _, edge := range g.Neighbours(u) {
		if !visited[edge.To] {
			dfs_top_sc(g, visited, stack, edge.To)
		}
	}

	*stack = append(*stack, u)
}
