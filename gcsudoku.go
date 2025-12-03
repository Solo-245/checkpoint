/*package main

import (
	"fmt"
	"os"
)

func printError() {
	fmt.Println("Error")
	os.Exit(0)
}

func isValid(grid [9][9]rune, row int, col int, num rune) bool {
	for i := 0; i < 9; i++ {
		if grid[row][i] == num || grid[i][col] == num {
			return false
		}
	}
	startRow := (row / 3) * 3
	startCol := (col / 3) * 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[startRow+i][startCol+j] == num {
				return false
			}
		}
	}
	return true
}

func isInitialValid(grid [9][9]rune) bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if grid[row][col] != '.' {
				num := grid[row][col]
				grid[row][col] = '.'
				if !isValid(grid, row, col, num) {
					return false
				}
				grid[row][col] = num
			}
		}
	}
	return true
}

func solve(grid *[9][9]rune) bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if grid[row][col] == '.' {
				for num := '1'; num <= '9'; num++ {
					if isValid(*grid, row, col, num) {
						grid[row][col] = num
						if solve(grid) {
							return true
						}
						grid[row][col] = '.'
					}
				}
				return false
			}
		}
	}
	return true
}

func main() {
	if len(os.Args) != 10 {
		printError()
	}

	var grid [9][9]rune

	for i := 0; i < 9; i++ {
		arg := os.Args[i+1]
		if len(arg) != 9 {
			printError()
		}
		for j := 0; j < 9; j++ {
			c := rune(arg[j])
			if (c < '1' || c > '9') && c != '.' {
				printError()
			}
			grid[i][j] = c
		}
	}

	if !isInitialValid(grid) {
		printError()
	}

	if !solve(&grid) {
		printError()
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%c", grid[i][j])
			if j != 8 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

}
