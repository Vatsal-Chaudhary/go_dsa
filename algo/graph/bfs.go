package graph

import (
	dsgraph "go_dsa/ds/graph"
)

func bfs_from(g dsgraph.Graph, src int, visited map[int]bool) []int {
	queue := []int{src}
	visited[src] = true
	result := []int{}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		// fmt.Println(node)
		result = append(result, node)

		for _, edge := range g.Neighbours(node) {
			if !visited[edge.To] {
				visited[edge.To] = true
				queue = append(queue, edge.To)
			}
		}
	}

	return result
}

// BFS from single source
func BFS(g dsgraph.Graph, src int) []int {
	visited := make(map[int]bool)
	return bfs_from(g, src, visited)
}

// BFS over full graph, handles disconnected source
func BFSDisconnected(g dsgraph.Graph) [][]int {
	visited := make(map[int]bool)
	result := [][]int{}

	for _, v := range g.Vertices() {
		if !visited[v] {
			r := bfs_from(g, v, visited)
			result = append(result, r)
		}
	}

	return result
}

// Counting connected Components in a connected graph
func BFSCount(g dsgraph.Graph) int {
	visited := make(map[int]bool)
	count := 0

	for _, v := range g.Vertices() {
		if !visited[v] {
			_ = bfs_from(g, v, visited)
			count++
		}
	}

	return count
}
