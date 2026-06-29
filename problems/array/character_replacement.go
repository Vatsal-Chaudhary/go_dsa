package array

func characterReplacement(s string, k int) int {
	l := len(s)
	if l == 0 {
		return 0
	}

	left := 0
	answer := 0

	freq := make([]int, 26)

	maxFreq := 0

	for right := 0; right < l; right++ {
		freq[s[right]-'A']++

		maxFreq = max(maxFreq, freq[s[right]-'A'])

		for ((right - left + 1) - maxFreq) > k {
			freq[s[left]-'A']--

			left++
		}

		answer = max(answer, right-left+1)
	}

	return answer
}
