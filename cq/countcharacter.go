//package main

import "fmt"

func Countcharacter(str string, c rune) int {
	count := 0
	for _, r := range str {
		if r == c {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(Countcharacter("Hello World", 'l'))
	fmt.Println(Countcharacter("5 balloons", 5))
	fmt.Println(Countcharacter("   ", ' '))
	fmt.Println(Countcharacter("The 7 deadly sins", '7'))
}
