package string

func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	count := make(map[rune]int)

	for _, ch := range s1 {
		count[ch]++
	}

	for _, ch := range s2 {
		count[ch]--
	}

	for _, value := range count {
		if value != 0 {
			return false
		}
	}

	return true
}
