package printnbr

import "github.com/01-edu/z01"

func PrintNbr(nb int) {
	abcd := 1
	if nb == 0 {
		z01.PrintRune('0')
		return
	}
	if nb < 0 {
		abcd *= -1
		z01.PrintRune('-')
	}
	PrintMine(nb, abcd)
}

func PrintMine(nb, abcd int) {
	if nb == 0 {
		return
	}
	digit := (nb % 10) * abcd
	_nb := nb / 10
	PrintMine(_nb, abcd)
	z01.PrintRune(rune(digit + '0'))
}
