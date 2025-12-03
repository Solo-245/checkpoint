package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	s1 := []rune(os.Args[1])
	s2 := os.Args[2]

	idx := 0

	for _, char := range s2 {
		if idx < len(s1) {
			if char == s1[idx] {
				idx++
			}
		}
	}

	if idx == len(s1) {
		for _, char := range s1 {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}

/*
$ go run . 123 123
123
$ go run . faya fgvvfdxcacpolhyghbreda
faya
$ go run . faya fgvvfdxcacpolhyghbred
$ go run . error rrerrrfiiljdfxjyuifrrvcoojh
$ go run . "quarante deux" "qfqfsudf arzgsayns tsregfdgs sjytdekuoixq "
quarante deux
$ go run .
$
*/
