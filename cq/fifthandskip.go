//package main

import "fmt"

func fifthAndSkip(str string) string {
	if str == "" {
		return "\n"
	}
	runes := []rune(str)

	if len(runes) < 5 {
		return "Invalid Input\n"
	}
	var result []rune
	count := 0
	skip := false
	for _, r := range runes {
		if r == ' ' {
			continue
		}
		if skip {
			skip = false
			continue
		}
		result = append(result, r)
		count++

		if count == 5 {
			result = append(result, ' ')
			count = 0
			skip = true
		}
	}

	if len(result) > 0 && result[len(result)-1] == ' ' {
		result = result[:len(result)-1]
	}
	return string(result)
}

func main() {
	fmt.Println(fifthAndSkip("abcdefghijklmnopqrstuwxyz"))
	fmt.Println(fifthAndSkip("This is a short sentence"))
	fmt.Println(fifthAndSkip("1234"))
}
