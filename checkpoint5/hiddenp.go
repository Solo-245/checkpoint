/*package main

import (
	"os"
	"strings"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	s1 := os.Args[1]
	s2 := os.Args[2]

	currentPos := 0

	for _, r := range s1 {
		foundAt := strings.Index(s2[currentPos:], string(r))

		if foundAt == -1 {
			z01.PrintRune('0')
			z01.PrintRune('\n')
			return
		}

		currentPos += foundAt + 1
	}

	z01.PrintRune('1')
	z01.PrintRune('\n')
}*/

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	s1 := os.Args[1]
	s2 := os.Args[2]

	if len(s1) == 0 {
		z01.PrintRune('1')
		z01.PrintRune('\n')
		return
	}

	i := 0
	for j := 0; j < len(s2); j++ {
		if s1[i] == s2[j] {
			i++
		}
		if i == len(s1) {
			z01.PrintRune('1')
			z01.PrintRune('\n')
			return
		}
	}

	z01.PrintRune('0')
	z01.PrintRune('\n')
}
