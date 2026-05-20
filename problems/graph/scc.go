package graph

import (
	dsgraph "go_dsa/ds/graph"
)

func TarjanSCC(g dsgraph.Graph) [][]int {
	visited := make(map[int]bool)
	disc := make(map[int]int)
	low := make(map[int]int)
	stack := []int{}
	onStack := make(map[int]bool)
	timer := 0
	result := [][]int{}

	for _, u := range g.Vertices() {
		if !visited[u] {
			dfs_scc(g, u, visited, disc, low, &stack, &result, &timer, onStack)
		}
	}

	return result
}

func dfs_scc(
	g dsgraph.Graph,
	u int,
	visited map[int]bool,
	disc map[int]int,
	low map[int]int,
	stack *[]int,
	result *[][]int,
	timer *int,
	onStack map[int]bool,
) {
	visited[u] = true

	*timer++
	low[u] = *timer
	disc[u] = *timer
	*stack = append(*stack, u)
	onStack[u] = true


	for _, edge := range g.Neighbours(u) {
		v := edge.To
		if !visited[edge.To] {

			dfs_scc(g, v, visited, disc, low, stack, result, timer, onStack)

			low[u] = min(low[u], low[v])

		} else if onStack[v] {
			low[u] = min(low[u], disc[v])
		}
	}

	if low[u] == disc[u] {
		scc := []int{}

		for {
					n := (*stack)[len(*stack)-1]
					*stack = (*stack)[:len(*stack)-1]
					onStack[n] = false

					scc = append(scc, n)

					if n == u {
						break
					}
				}

				*result = append(*result, scc)
	}
}




func canFinish(numCourses int, prerequisites [][]int) bool {
	visited := make([]bool, numCourses)
	recSt := make([]bool, numCourses)

	graph := make([][]int, numCourses)

	for _, p := range prerequisites {
		course := p[0]
		prereq := p[1]

		graph[prereq] = append(graph[prereq], course)
	}

	for i := 0; i < numCourses; i++ {
		if !visited[i] {
			if dfs_course(graph, i, visited, recSt) {
				return false
			}
		}
	}

	return true
}

func dfs_course(graph [][]int, i int, visited []bool, recSt []bool) bool {
	visited[i] = true
	recSt[i] = true

	for _, v := range graph[i] {
		if !visited[v] {
			if dfs_course(graph, v, visited, recSt) {
				return true
			}
		} else if recSt[v] {
			return true
		}
	}

	recSt[i] = false

	return false
} 
