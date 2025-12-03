//package main

import "fmt"

func Printifnot(str string) string {
	if len(str) < 3 {
		return "G\n"
	}
	return "Invalid Input\n"
}

func main() {
	fmt.Print(Printifnot("abcdefz"))
	fmt.Print(Printifnot("abc"))
	fmt.Print(Printifnot(""))
	fmt.Print(Printifnot("14"))
}
