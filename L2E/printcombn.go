package student

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}

	combination := make([]rune, n)
	generateCombinations(combination, 0, '0')
	z01.PrintRune('\n')
}

func generateCombinations(combination []rune, index int, start rune) {
	if index == len(combination) {
		for _, digit := range combination {
			z01.PrintRune(digit)
		}

		if !isLastCombination(combination) {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
		return
	}

	for ch := start; ch <= '9'; ch++ {
		combination[index] = ch
		generateCombinations(combination, index+1, ch+1)
	}
}

func isLastCombination(combination []rune) bool {
	n := len(combination)
	for i := 0; i < n; i++ {
		if combination[i] != rune(i)+('9'-rune(n)+1) {
			return false
		}
	}
	return true
}
