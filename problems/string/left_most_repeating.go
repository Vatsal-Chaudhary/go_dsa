package string

func leftMostRepeatingChar_1(str string) int {
	if len(str) == 0 {
		return -1
	}

	count := make(map[rune]int)

	for _, c := range str {
		count[c]++
	}

	for i, c := range str {
		if count[c] > 1 {
			return i
		}
	}

	return -1
}

func leftMostRepeatingChar(str string) int {
	if len(str) == 0 {
		return -1
	}

	visited := make(map[byte]bool)
	res := -1

	for i := len(str) - 1; i >= 0; i-- {
		if visited[str[i]] {
			res = i
		} else {
			visited[str[i]] = true
		}
	}

	return res
}
