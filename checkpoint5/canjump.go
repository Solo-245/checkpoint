package main

import (
	"fmt"
)

func CanJump(data []uint) bool {
	if len(data) == 0 {
		return false
	}
	if len(data) == 1 {
		return true
	}

	index := 0
	lastIndex := len(data) - 1

	for index < lastIndex {
		jump := int(data[index])

		if jump == 0 {
			return false
		}

		index += jump
	}

	return index == lastIndex
}

func main() {
	input1 := []uint{2, 3, 1, 1, 4}
	fmt.Println(CanJump(input1))

	input2 := []uint{3, 2, 1, 0, 4}
	fmt.Println(CanJump(input2))

	input3 := []uint{0}
	fmt.Println(CanJump(input3))
}
