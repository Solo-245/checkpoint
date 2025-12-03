package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	var count [10]int

	for n > 0 {
		d := n % 10
		count[d]++
		n /= 10
	}

	for digit := 0; digit <= 9; digit++ {
		for count[digit] > 0 {
			z01.PrintRune(rune('0' + digit))
			count[digit]--
		}
	}
}
