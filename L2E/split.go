package piscine

func Split(s, sep string) []string {
	var result []string
	sepLen := len(sep)
	if sepLen == 0 {
		return []string{s}
	}

	start := 0

	for i := 0; i+sepLen <= len(s); i++ {
		if s[i:i+sepLen] == sep {
			result = append(result, s[start:i])
			start = i + sepLen
			i = i + sepLen - 1
		}
	}

	result = append(result, s[start:])

	return result
}
