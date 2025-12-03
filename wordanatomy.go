package main

import "fmt"

func WordAnatomy(s string, pre []string, suf []string) string {
	mp, ms := "", ""

	for _, v := range pre {
		if hasPrefix(s, v) {
			mp = v
			break
		}
	}
	for _, v := range suf {
		if hasSuffix(s, v) {
			ms = v
			break
		}
	}

	return "prefix: " + mp + "," + " suffix: " + ms
}

func hasPrefix(s, sub string) bool {
	if len(sub) > len(s) {
		return false
	}
	return s[:len(sub)] == sub
}

func hasSuffix(s, sub string) bool {
	if len(sub) > len(s) {
		return false
	}
	return s[len(s)-len(sub):] == sub
}

func main() {
	fmt.Println(WordAnatomy("unbelievable", []string{"un", "dis"}, []string{"able", "ing"}))

	fmt.Println(WordAnatomy("understand", []string{"un", "under"}, []string{"stand", "and"}))

	fmt.Println(WordAnatomy("hello", []string{"x", "y"}, []string{"z"}))

	fmt.Println(WordAnatomy("hi", []string{"hello"}, []string{"world"}))

	fmt.Println(WordAnatomy("test", []string{"test"}, []string{"test"}))
}
