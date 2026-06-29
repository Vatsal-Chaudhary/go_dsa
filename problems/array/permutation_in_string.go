package array

func checkInclusion(s1, s2 string) bool {
	l1 := len(s1)
	l2 := len(s2)

	if l1 > l2 {
		return false
	}

	var freq [26]int
	var window [26]int

	for i := 0; i < l1; i++ {
		freq[s1[i]-'a']++
	}

	for right := 0; right < l2; right++ {
		window[s2[right]-'a']++

		if right >= l1 {
			left := s2[right-l1]
			window[left-'a']--
		}

		if window == freq {
			return true
		}
	}

	return false
}
