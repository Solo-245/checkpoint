package main

import "fmt"

/*func iterativefunction(index int) int {
	result := 0

	for i := 0; i < index+1; i++ {
		result = result + i
	}
	return result
}

func recursivecalculation(index int) int {
	if index == 1 {
		return 1
	}
	if index > 1 {
		return index * recursivecalculation(index-1)
	}
	return 0
}

func fact(nb int) int {
	if nb == 1 {
		return 1
	}
	if nb > 1 {
		return nb * fact(nb-1)
	}
	return 0
}

func main() {
	result1 := 0 + 1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9 + 10 + 11 + 12 + 13
	fmt.Println(result1)

	result2 := iterativefunction(14)
	fmt.Printf("the result of the index is: %v \n", result2)
	result3 := recursivecalculation(5)
	fmt.Printf("the result of the recursive calculation is: %v \n", result3)
	result4 := fact(4)
	fmt.Printf("the result of the fact is: %v \n", result4)
}*/

func Fibonacci(n int) int {
	if n < 2 {
		return 1
	}
	return Fibonacci((n - 1)) + Fibonacci(n-2)
}
func main() {
	fmt.Println(Fibonacci(8))
}
