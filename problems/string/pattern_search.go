package string

func SearchPattern(txt, pattern string) []int {
	if txt == "" || pattern == "" {
		return []int{}
	}

	runesP := []rune(pattern)
	runesT := []rune(txt)

	m := len(runesP)
	n := len(runesT)
	
	res := []int{}

	for i := 0; i <= n - m; i++ {
		j := 0
		for j = 0; j < m; j++ {
			if runesP[j] != runesT[j + i] {
				break
			}
		}
		if j == m {
			res = append(res, i)
		}
	}

	return res
}
