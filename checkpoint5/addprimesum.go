package main

import (
	"os"

	"github.com/01-edu/z01"
)

func basicAtoi(s string) int {
	n := 0
	for _, c := range s {
		if c < '0' || c > '9' {
			return 0
		}
		n = n*10 + int(c-'0')
	}
	return n
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func printNbr(n int) {
	if n >= 10 {
		printNbr(n / 10)
	}
	z01.PrintRune(rune(n%10 + '0'))
}

func main() {
	if len(os.Args) != 2 {
		printNbr(0)
		z01.PrintRune('\n')
		return
	}

	num := basicAtoi(os.Args[1])

	if num <= 0 {
		printNbr(0)
		z01.PrintRune('\n')
		return
	}

	sum := 0
	for i := 2; i <= num; i++ {
		if isPrime(i) {
			sum += i
		}
	}

	printNbr(sum)
	z01.PrintRune('\n')
}

/*package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func printNbr(n int) {
	if n >= 10 {
		printNbr(n / 10)
	}
	z01.PrintRune(rune(n%10 + '0'))
}

func main() {
	if len(os.Args) != 2 {
		printNbr(0)
		z01.PrintRune('\n')
		return
	}

	num, err := strconv.Atoi(os.Args[1])

	if err != nil || num <= 0 {
		printNbr(0)
		z01.PrintRune('\n')
		return
	}

	sum := 0
	for i := 2; i <= num; i++ {
		if isPrime(i) {
			sum += i
		}
	}

	printNbr(sum)
	z01.PrintRune('\n')
}*\
