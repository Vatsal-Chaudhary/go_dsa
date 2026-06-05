package string

func LongestDistinctSubstring_naive(str string) int {
	runes := []rune(str)
	n := len(runes)
	if n == 0 {
		return 0
	}

	res := 0

	for i := 0; i < n; i++ {
		visited := make(map[rune]bool)

		for j := i; j < n; j++ {
			if visited[runes[j]] {
				break
			} else {
				res = max(res, j - i + 1)
				visited[runes[j]] = true
			}
		}
	}

	return res
}

func LongestDistinctSubstring(str string) int {
	runes := []rune(str)
	n := len(runes)
	if n == 0 {
		return 0
	}

	res := 0

	prev := make(map[rune]int)

	for i := 0; i < n; i++ {
		prev[runes[i]] = -1
	}

	j := 0
	for i := 0; i < n; i++ {
		j = max(j, prev[runes[i]] + 1)

		maxEnd := i - j + 1

		res = max(res, maxEnd)

		prev[runes[i]] = i
	}

	return res
}
