//package main

import "fmt"

func Printif(str string) string {
	if len(str) == 0 || len(str) >= 3 {
		return "G\n"
	}
	return "Invalid Input\n"
}

func main() {
	fmt.Print(Printif("abcdefz"))
	fmt.Print(Printif("abc"))
	fmt.Print(Printif(""))
	fmt.Print(Printif("14"))
}
