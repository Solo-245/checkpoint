package student

func BasicAtoi(s string) int {
	result := 0

	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			result = result*10 + int(ch-'0')
		}
	}
	return result
}
