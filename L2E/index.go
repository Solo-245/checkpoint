package piscine

func Index(s string, toFind string) int {
	slength := len(s)
	toFindLen := len(toFind)

	if toFindLen == 0 {
		return 0
	}

	for i := 0; i <= slength-toFindLen; i++ {
		j := 0
		for j < toFindLen && s[i+j] == toFind[j] {
			j++
		}
		if j == toFindLen {
			return i
		}
	}
	return -1
}
