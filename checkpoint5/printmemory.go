package main

import "github.com/01-edu/z01"

func PrintMemory(arr [10]byte) {
	for i, b := range arr {
		printHex(b)
		if (i+1)%4 == 0 {
			z01.PrintRune('\n')
		} else if i != len(arr)-1 {
			z01.PrintRune(' ')
		} else {
			z01.PrintRune('\n')
		}
	}

	for _, b := range arr {
		if b >= 32 && b <= 126 {
			z01.PrintRune(rune(b))
		} else {
			z01.PrintRune('.')
		}
	}
	z01.PrintRune('\n')
}

func printHex(n byte) {
	base := "0123456789abcdef"
	z01.PrintRune(rune(base[n/16]))
	z01.PrintRune(rune(base[n%16]))
}

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*', 100, 200})
}
