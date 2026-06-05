package string

func isSubsequence(subseq, str string) bool {
	if len(str) == 0 && len(subseq) != 0 {
		return false
	}

	j := 0
	for i := 0; i < len(str) && j < len(subseq); i++ {
		if str[i] == subseq[j] {
			j++
		}
	}

	return j == len(subseq)
}
