package main

import (
	"os"
	"strings"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		z01.PrintRune('\n')
		return
	}

	str1 := os.Args[1]
	str2 := os.Args[2]

	printed := ""

	for _, char := range str1 {
		if !strings.ContainsRune(printed, char) {
			z01.PrintRune(char)
			printed += string(char)
		}
	}

	for _, char := range str2 {
		if !strings.ContainsRune(printed, char) {
			z01.PrintRune(char)
			printed += string(char)
		}
	}

	z01.PrintRune('\n')
}

/*
$ go run . zpadinton paqefwtdjetyiytjneytjoeyjnejeyj | cat -e
zpadintoqefwjy$
$
$ go run . ddf6vewg64f gtwthgdwthdwfteewhrtag6h4ffdhsd | cat -e
df6vewg4thras$
$
$ go run . rien "cette phrase ne cache rien" | cat -e
rienct phas$
$
$ go run . | cat -e
$
$ go run . rien | cat -e
$
*/
