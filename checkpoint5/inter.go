/*package main

import (
	"os"
	"strings"


	"github.com/01-edu/z01"
)

/*func main() {
	if len(os.Args) != 3 {
		return
	}

	s1 := os.Args[1]
	s2 := os.Args[2]

	seen := ""

	for _, char := range s1 {
		if strings.IndexRune(s2, char) >= 0 && strings.IndexRune(seen, char) == -1 {
			z01.PrintRune(char)
			seen += string(char)
		}
	}

	z01.PrintRune('\n')
}

/*
$ go run . "padinton" "paqefwtdjetyiytjneytjoeyjnejeyj"
padinto
$ go run . ddf6vewg64f  twthgdwthdwfteewhrtag6h4ffdhsd
df6ewg4
$
*/

package main

import (
	"github.com/01-edu/z01"
	"os"
)

func contains(s string, r rune) bool {
	for _, c := range s {
		if c == r {
			return true
		}
	}
	return false
}

func main() {
	if len(os.Args) != 3 {
		return
	}

	s1 := os.Args[1]
	s2 := os.Args[2]
	seen := ""

	for _, char := range s1 {
		if contains(s2, char) && !contains(seen, char) {
			z01.PrintRune(char)
			seen += string(char)
		}
	}

	z01.PrintRune('\n')
}
