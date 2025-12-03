//package main

import (
	"fmt"
)

func NotDecimal(dec string) string {
	if dec == "" {
		return "\n"
	}

	dotIndex := -1
	dotCount := 0
	startIdx := 0

	for i, char := range dec {
		if char == '.' {
			dotIndex = i
			dotCount++
		} else if (char == '-' || char == '+') && i == 0 {
			startIdx = 1
		} else if char < '0' || char > '9' {
			return dec + "\n"
		}
	}

	if dotCount != 1 {
		return dec + "\n"
	}

	afterDot := dec[dotIndex+1:]
	if afterDot == "" || afterDot == "0" {
		return dec + "\n"
	}

	sign := ""
	if startIdx == 1 {
		sign = string(dec[0])
	}

	beforeDot := dec[startIdx:dotIndex]
	combined := beforeDot + afterDot

	nonZeroIndex := 0
	for nonZeroIndex < len(combined) && combined[nonZeroIndex] == '0' {
		nonZeroIndex++
	}

	res := combined[nonZeroIndex:]

	if res == "" {
		res = "0"
	}

	return sign + res + "\n"
}

func main() {
	fmt.Print(NotDecimal("0.1"))
	fmt.Print(NotDecimal("174.2"))
	fmt.Print(NotDecimal("0.1255"))
	fmt.Print(NotDecimal("1.20525856"))
	fmt.Print(NotDecimal("-0.0f00d00"))
	fmt.Print(NotDecimal(""))
	fmt.Print(NotDecimal("-19.525856"))
	fmt.Print(NotDecimal("1952"))
}
