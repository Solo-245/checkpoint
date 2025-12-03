/*
package main
import "fmt"

	func main() {
		sentence := []byte("Hello")
		counter := 0
		//nums := []int{1, 2, 3, 4, 5}
		for index, letter := range sentence {
			counter++
			fmt.Println(sentence, index, letter)
			fmt.Print(counter)
		}
	}

package main

import "fmt"

	func CheckNumber(arg string) bool {
		for _, r := range arg {
			if r >= '0' && r <= '9' {
				return true
			}
		}
		return false
	}

	func main() {
		fmt.Println(CheckNumber("Hello"))
		fmt.Println(CheckNumber("Hello1"))
	}

package main
import "fmt"

	func CountAlpha(s string) int {
		count := 0
		for _, r := range s {
			if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' {
				count++
			}
		}
		return count
	}

	func main() {
		fmt.Println(CountAlpha("Hello World"))
		fmt.Println(CountAlpha("Hel   78u ib 89"))
	}

package main

import "fmt"

	func CountCharacter(str string, c rune) int {
		count := 0
		for _, r := range str {
			if r == c {
				count++
			}
		}
		return count
	}

	func main() {
		fmt.Println(CountCharacter("Hello World", 'l'))
		fmt.Println(CountCharacter("5 ballons", 5))
	}

package main

import "fmt"

	func Printif(str string) string {
		if len(str) == 0 || len(str) >= 3 {
			return "G\n"
		}
		return "Invalid Input\n"
	}

	func main() {
		fmt.Print(Printif("abcdefz"))
		fmt.Print(Printif("14"))
	}

package main

import "fmt"

	func PrintIfNot(str string) string {
		if len(str) < 3 {
			return "G\n"
		}
		return "Invalid Input\n"
	}

	func main() {
		fmt.Print(PrintIfNot("abcdefg"))
		fmt.Print(PrintIfNot("14d"))
	}

package main

import "fmt"

	func RectPerimeter(w, h int) int {
		if w < 0 || h < 0 {
			return -1
		} else {
			return 2 * (w + h)
		}
	}

	func main() {
		fmt.Println(RectPerimeter(2, 3))
		fmt.Println(RectPerimeter(10, 2))
	}

package main

import "fmt"

	func retainfirsthalf(str string) string {
		n := len(str)
		if n <= 1 {
			return str
		}
		half := n / 2
		return str[:half]

}

package main

import (
	"fmt"
	"strings"
)

func RetainFirstHalf(str string) string {
	if len(str) == 0 {
		return ""
	} else if len(str) == 1 {
		return (str)
	} else {
		var res strings.Builder
		i := 0
		for i = 0; i < int(len(str)/2); i++ {
			res.WriteRune(rune(str[i]))
		}
		return res.String()
	}
}

func main() {
	fmt.Println(RetainFirstHalf("This is the 1st halfThis is the 2nd half"))
	fmt.Println(RetainFirstHalf("A"))
	fmt.Println(RetainFirstHalf(""))
	fmt.Println(RetainFirstHalf("Hello World"))
}*/
