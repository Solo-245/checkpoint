package piscine

func IsPrintable(s string) bool {
	runes := []rune(s)
	for i := range runes {
		if CheckPrint(runes[i]) == false {
			return false
		}
	}
	return true
}

func CheckPrint(x rune) bool {
	if x >= 32 && x <= 126 {
		return true
	}
	return false
}
