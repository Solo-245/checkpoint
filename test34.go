//package main

import "fmt"

func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}
	runes := []rune(s)
	length := len(runes)
	for i := 0; i < length; i++ {
		r := runes[i]
		isUpper := r >= 'A' && r <= 'Z'
		isLower := r >= 'a' && r <= 'z'
		if !isUpper && !isLower {
			return s
		}
		if i == length-1 && isUpper {
			return s
		}
		if i < length-1 {
			next := runes[i+1]
			nextIsUpper := next >= 'A' && next <= 'Z'
			if isUpper && nextIsUpper {
				return s
			}
		}
	}
	var result []rune
	for i := 0; i < length; i++ {
		r := runes[i]
		isUpper := r >= 'A' && r <= 'Z'
		if isUpper {
			if i > 0 {
				result = append(result, '_')
			}
			result = append(result, r)

		} else {
			result = append(result, r)
		}
	}
	return string(result)
}

func main() {
	fmt.Println(CamelToSnakeCase("camelCase"))
}
