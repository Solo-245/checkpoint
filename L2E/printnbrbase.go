package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if !isValidBase(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	baseLen := len(base)
	if nbr < 0 {
		z01.PrintRune('-')
		if nbr == -9223372036854775808 {
			printNbrBaseRecursive(-(nbr / baseLen), base)
			z01.PrintRune(rune(base[-(nbr % baseLen)]))
			return
		}
		nbr = -nbr
	}

	printNbrBaseRecursive(nbr, base)
}

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}

	for i, r1 := range base {
		if r1 == '+' || r1 == '-' {
			return false
		}
		for j, r2 := range base {
			if i != j && r1 == r2 {
				return false
			}
		}
	}
	return true
}

func printNbrBaseRecursive(nbr int, base string) {
	baseLen := len(base)
	if nbr >= baseLen {
		printNbrBaseRecursive(nbr/baseLen, base)
	}
	z01.PrintRune(rune(base[nbr%baseLen]))
}
