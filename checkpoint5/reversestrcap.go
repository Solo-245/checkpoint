/*package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		return
	}

	for _, s := range args {
		runes := []rune(s)
		length := len(runes)

		for i, r := range runes {
			if r == ' ' {
				z01.PrintRune(' ')
				continue
			}

			isLast := false
			if i == length-1 {
				isLast = true
			} else if runes[i+1] == ' ' {
				isLast = true
			}

			if isLast {
				z01.PrintRune(toUpper(r))
			} else {
				z01.PrintRune(toLower(r))
			}
		}
		z01.PrintRune('\n')
	}
}

func toUpper(r rune) rune {
	if r >= 'a' && r <= 'z' {
		return r - 32
	}
	return r
}

func toLower(r rune) rune {
	if r >= 'A' && r <= 'Z' {
		return r + 32
	}
	return r
}

/*
$ go run . "First SMALL TesT" | cat -e
firsT smalL tesT$
$ go run . "SEconD Test IS a LItTLE EasIEr" "bEwaRe IT'S NoT HARd WhEN " " Go a dernier 0123456789 for the road e" | cat -e
seconD tesT iS A littlE easieR$
bewarE it'S noT harD wheN $
 gO A dernieR 0123456789 foR thE roaD E$
$ go run .
$
*/

package main

import (
	"fmt"
	"os"
	"unicode"
)

func main() {
	for _, arg := range os.Args[1:] {
		arg := []rune(arg)
		for j, r := range arg {
			if j+1 == len(arg) || arg[j+1] == ' ' {
				arg[j] = unicode.ToUpper(r)
			} else {
				arg[j] = unicode.ToLower(r)
			}
		}
		fmt.Println(string(arg))
	}
}
