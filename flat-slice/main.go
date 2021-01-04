package main

import (
	"fmt"
	"math"
)

var input = [][]int{
	{1, 2, 3, 4, 5},
	{5, 6, 7, 8, 6},
	{9, 10, 11, 12, 7},
	{13, 14, 15, 16, 8},
	{55, 67, 66, 10, 13},
}

func main() {
	if !isValidDimension(input) {
		fmt.Println("Number of columns and rows must be equal.")
		return
	}

	fmt.Println(FlattenSlice(input))
}

func FlattenSlice(s [][]int) []int {
	var res []int
	length := int(math.Ceil(float64(len(s)) / float64(2)))

	for i := 0; i < length; i++ {
		res = append(res, goForward(s, i)...)
		res = append(res, goBackward(s, i)...)
	}

	return res
}

func isValidDimension(s [][]int) bool {
	for _, v := range s {
		if len(v) != len(s) {
			return false
		}
	}

	return true
}

func goForward(s [][]int, startRow int) []int {
	res := s[startRow][startRow : len(s)-startRow]

	for i := startRow + 1; i < len(s)-startRow; i++ {
		res = append(res, s[i][len(s)-(startRow+1)])
	}

	return res
}

func goBackward(s [][]int, startRow int) []int {
	startRowBack := startRow + 1
	res := reverse(s[len(s)-startRowBack][startRow : len(s)-startRowBack])

	for i := len(s) - (startRowBack + 1); i >= startRowBack; i-- {
		res = append(res, s[i][startRow])
	}

	return res
}

func reverse(s []int) []int {
	var res = make([]int, len(s))

	for i, v := range s {
		res[len(s)-(i+1)] = v
	}

	return res[:]
}
