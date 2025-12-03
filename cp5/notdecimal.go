package main

import (
	"fmt"
	"math"
	"strconv"
)

func NotDecimal(dec string) string {
	j := -1
	n := 0
	if len(dec) == 0 {
		return "\n"
	}
	for i := 0; i < len(dec); i++ {
		if j == -1 && dec[i] == '.' {
			j++
		} else if j == 0 {
			n++
		}
	}
	s, err := strconv.ParseFloat(dec, 64)
	if err == nil {
		return strconv.Itoa(int(s*math.Pow(10, float64(n)))) + "\n"
	}
	return (dec + "\n")
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

/*
$ go run .  | cat -e
1$
1742$
1255$
120525856$
-0.0f00d00$
$
-19525856$
1952$
*/
