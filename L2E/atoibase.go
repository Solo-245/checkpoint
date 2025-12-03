package piscine

func AtoiBase(s string, base string) int {
	if !checkValid(base) {
		return 0
	}
	checkValid(base)
	baseRn := []rune(base)
	numConv := []rune(s)
	baseLen := len(base)
	numIndexStart := len(numConv)

	result := 0
	index := numIndexStart - 1
	for i := 0; i < len(numConv); i++ {
		for j := 0; j < len(baseRn); j++ {
			if numConv[i] == baseRn[j] {
				curr := j * calPower(int(baseLen), index)
				result += curr
				index--
			}
		}
	}
	return result
}

func calPower(nb, power int) int {
	if power < 0 {
		return 0
	}
	result := 1

	for i := 0; i < power; i++ {
		result *= nb
	}

	return result
}

func checkValid(base string) bool {
	if len(base) < 2 {
		return false
	}

	baseRune := []rune(base)

	for i := 0; i < len(baseRune); i++ {
		if baseRune[i] == '+' || baseRune[i] == '-' {
			return false
		}

		for x := i + 1; x < len(baseRune); x++ {
			if baseRune[i] == baseRune[x] {
				return false
			}
		}
	}
	return true
}
