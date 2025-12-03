package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for a := '0'; a <= '9'; a++ {
		for b := '0'; b <= '9'; b++ {
			// Calculate the first two-digit number
			i := (a-'0')*10 + (b - '0')

			for c := '0'; c <= '9'; c++ {
				for d := '0'; d <= '9'; d++ {
					// Calculate the second two-digit number
					j := (c-'0')*10 + (d - '0')

					// Error was here: The condition should be i < j to ensure the first number is strictly less than the second number
					if i < j {
						z01.PrintRune(a)
						z01.PrintRune(b)
						z01.PrintRune(' ')
						z01.PrintRune(c)
						z01.PrintRune(d)

						// The condition for the final comma also needs to check the last possible pair: 98 99
						if i != 98 || j != 99 {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}

func main() {
	PrintComb2()
}
