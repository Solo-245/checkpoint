/*package main

import (
	"fmt"
	"os"
)

var field [9][9]int

func canPut(col int, row int, value int) bool {
	startRow, startCol := (row/3)*3, (col/3)*3

	for i := 0; i < 9; i++ {
		if field[row][i] == value || field[i][col] == value || field[startRow+i/3][startCol+i%3] == value {
			return false
		}
	}
	return true
}

func solve(col int, row int) bool {
	if row == 9 {
		return true
	}

	nextCol, nextRow := col+1, row
	if nextCol == 9 {
		nextCol, nextRow = 0, row+1
	}

	if field[row][col] != 0 {
		return solve(nextCol, nextRow)
	}

	for value := 1; value <= 9; value++ {
		if canPut(col, row, value) {
			field[row][col] = value
			if solve(nextCol, nextRow) {
				return true
			}
			field[row][col] = 0
		}
	}
	return false
}

func main() {
	if len(os.Args) != 10 {
		fmt.Println("Error")
		return
	}

	for row := 0; row < 9; row++ {
		line := os.Args[row+1]
		if len(line) != 9 {
			fmt.Println("Error")
			return
		}
		for col := 0; col < 9; col++ {
			char := line[col]
			if char == '.' {
				field[row][col] = 0
			} else if char >= '1' && char <= '9' {
				field[row][col] = int(char - '0')
			} else {
				fmt.Println("Error")
				return
			}
		}
	}

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			value := field[row][col]
			if value != 0 {
				field[row][col] = 0
				if !canPut(col, row, value) {
					fmt.Println("Error")
					return
				}
				field[row][col] = value
			}
		}
	}

	if solve(0, 0) {
		for row := 0; row < 9; row++ {
			for col := 0; col < 9; col++ {
				if field[row][col] == 0 {
					fmt.Print(". ")
				} else {
					fmt.Printf("%d ", field[row][col])
				}
			}
			fmt.Println()
		}
	} else {
		fmt.Println("Error")
	}
}
