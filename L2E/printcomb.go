package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for a := '0'; a <= '9'; a++ {
		for b := '0'; b <= '9'; b++ {
			i := (a-'0')*10 + (b - '0')

			for c := '0'; c <= '9'; c++ {
				for d := '0'; d <= '9'; d++ {
					j := (c-'0')*10 + (d - '0')

					if b > a {
						z01.PrintRune(a)
						z01.PrintRune(b)
						z01.PrintRune(' ')
						z01.PrintRune(c)
						z01.PrintRune(d)

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
