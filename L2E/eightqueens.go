package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	var board [8]int
	solve(board, 0)
}

func solve(board [8]int, row int) {
	if row == 8 {
		printSolution(board)
		return
	}
	for col := 0; col < 8; col++ {
		if safe(board, row, col) {
			board[row] = col
			solve(board, row+1)
		}
	}
}

func safe(board [8]int, row, col int) bool {
	for i := 0; i < row; i++ {
		if board[i] == col ||
			abs(board[i]-col) == abs(i-row) {
			return false
		}
	}
	return true
}

func printSolution(board [8]int) {
	for i := 0; i < 8; i++ {
		z01.PrintRune(rune(board[i] + 1 + '0'))
	}
	z01.PrintRune('\n')
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
