//package main

import "fmt"

func RevConcatAlternate(slice1, slice2 []int) []int {
	len1 := len(slice1)
	len2 := len(slice2)

	totalLen := len1 + len2
	res := make([]int, 0, totalLen)

	if len1 > len2 {
		for i := len1 - 1; i >= len2; i-- {
			res = append(res, slice1[i])
		}
	} else if len2 > len1 {
		for i := len2 - 1; i >= len1; i-- {
			res = append(res, slice2[i])
		}
	}

	smallLen := len1
	if len2 < len1 {
		smallLen = len2
	}

	for i := smallLen - 1; i >= 0; i-- {
		res = append(res, slice1[i])
		res = append(res, slice2[i])
	}

	return res
}

func main() {

	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))

	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))

	fmt.Println(RevConcatAlternate([]int{1, 2, 3, 9, 8}, []int{4, 5}))

	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{}))

}
