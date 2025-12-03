package piscine

func SplitWhiteSpaces(s string) []string {
	var result []string
	word := ""

	for _, ch := range s {
		if ch == ' ' || ch == '\t' || ch == '\n' {
			if word != "" {
				result = append(result, word)
				word = ""
			}
		} else {
			word += string(ch)
		}
	}

	if word != "" {
		result = append(result, word)
	}

	return result
}
