package main

import "fmt"

func max(ints []int) int {
	highest := -9 * 7
	for _, v := range ints {
		if v > highest {
			highest = v
		}
	}
	return highest
}

func main() {
	var A [6][6]int
	for rowIndex, row := range A {
		for colIndex := range row {
			fmt.Scanf("%d", &A[rowIndex][colIndex])
		}
	}
	var hourglasses [16]int
	index := 0
	for rowIndex := 0; rowIndex < 4; rowIndex++ {
		for colIndex := 0; colIndex < 4; colIndex++ {
			hourglasses[index] = A[rowIndex][colIndex] + A[rowIndex][colIndex+1] + A[rowIndex][colIndex+2] + A[rowIndex+1][colIndex+1] + A[rowIndex+2][colIndex] + A[rowIndex+2][colIndex+1] + A[rowIndex+2][colIndex+2]
			index++
		}
	}
	fmt.Println(max(hourglasses[:]))
}
