package graph

import (
	"sort"
)

func normalizeComponents(components [][]int) [][]int {
	for _, c := range components {
		sort.Ints(c)
	}
	sort.Slice(components, func(i, j int) bool {
		return components[i][0] < components[j][0]
	})
	return components
}
