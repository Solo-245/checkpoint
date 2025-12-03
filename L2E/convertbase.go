package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	indexInBase := func(base string, r rune) int {
		for i, br := range base {
			if br == r {
				return i
			}
		}
		return -1
	}

	baseFromLen := len(baseFrom)
	value := 0
	for _, r := range nbr {
		digit := indexInBase(baseFrom, r)
		value = value*baseFromLen + digit
	}

	if value == 0 {
		return string(baseTo[0])
	}

	baseToLen := len(baseTo)
	result := ""
	for value > 0 {
		remainder := value % baseToLen
		result = string(baseTo[remainder]) + result
		value = value / baseToLen
	}

	return result
}
