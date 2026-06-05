package string

func validPalindrome(s string) bool {
	return check(s, 0, len(s) - 1, false)
}

func check(s string, l, r int, deleted bool) bool {

	for l < r {

		if s[l] != s[r] {
			if deleted {
				return false
			}

			return check(s, l+1, r, true) || check(s, l, r -1, true)
		}

		l++
		r--
	}

	return true
}
