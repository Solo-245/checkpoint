package main

/*import (
	"fmt"
	"strings"
)

func FirstWord(s string) string {
	words := strings.Fields(s)
	res := "\n"
	if len(words) > 0 {
		res = words[0] + res
	}
	return res
}

func main() {
	fmt.Print(FirstWord("hello there"))
	fmt.Print(FirstWord(""))
	fmt.Print(FirstWord("hello   .........  bye"))
}*/

/*
$ go run .
hello

hello
$
*/

import (
	"fmt"
	"strings"
)

func LastWord(s string) string {
	words := strings.Fields(s)
	if len(words) > 0 {
		return words[len(words)-1] + "\n"
	}
	return "\n"
}

func main() {
	fmt.Print(LastWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Print(LastWord(" lorem,ipsum "))
	fmt.Print(LastWord(" "))
}

/*
$ go run . | cat -e
not$
lorem,ipsum$
$
$
*/
