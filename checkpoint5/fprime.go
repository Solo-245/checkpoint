package main

/*import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	number, err := strconv.Atoi(os.Args[1])

	if err != nil || number <= 1 {
		return
	}

	divisor := 2

	for number > 1 {
		if number%divisor == 0 {
			fmt.Print(divisor)
			number = number / divisor

			if number > 1 {
				fmt.Print("*")
			}
		} else {
			divisor++
		}
	}
	fmt.Println()
}*/

/*
$ go run . 225225
3*3*5*5*7*11*13
$ go run . 8333325
3*3*5*5*7*11*13*37
$ go run . 9539
9539
$ go run . 804577
804577
$ go run . 42
2*3*7
$ go run . a
$ go run . 0
$ go run . 1
$
*/

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	s := os.Args[1]
	number := 0
	for _, c := range s {
		if c < '0' || c > '9' {
			return
		}
		number = number*10 + int(c-'0')
	}

	if number <= 1 {
		return
	}

	divisor := 2

	for number > 1 {
		if number%divisor == 0 {
			printNbr(divisor)
			number = number / divisor

			if number > 1 {
				z01.PrintRune('*')
			}
		} else {
			divisor++
		}
	}
	z01.PrintRune('\n')
}

func printNbr(n int) {
	if n >= 10 {
		printNbr(n / 10)
	}
	z01.PrintRune(rune(n%10 + '0'))
}
