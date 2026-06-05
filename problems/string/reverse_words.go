package string

func ReverseWords(str string) string {
	runes := []rune(str)

	var result string

	start := 0
	length := len(runes)

	for i := 0; i <= length; i++ {
		if i == length || runes[i] == ' ' {
			word := string(runes[start:i])
			if word != "" {
				if result == "" {
					result = word
				} else {
					result = word + " " + result
				}
			}
			start = i + 1
		}
	}

	return result
}
