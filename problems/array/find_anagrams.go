package array

func findAnagrams(s, p string) []int {
	if len(p) == 0 || len(p) > len(s) {
		return []int{}
	}

	var freq [26]int

	for i := 0; i < len(p); i++ {
		freq[p[i]-'a']++
	}

	var window [26]int
	res := []int{}

	for right := 0; right < len(s); right++ {
		window[s[right]-'a']++

		if right >= len(p) {
			left := right - len(p)
			window[s[left]-'a']--
		}

		if window == freq {
			res = append(res, right-len(p)+1)
		}
	}

	return res
}
