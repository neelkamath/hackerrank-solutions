package main

import "fmt"

type arr [6][6]int

func input() arr {
	var a arr
	for i, row := range a {
		for j := range row {
			fmt.Scanf("%d", &a[i][j])
		}
	}
	return a
}

func hourglass(array arr, row, col int) int {
	a := array[row][col]
	b := array[row][col+1]
	c := array[row][col+2]
	d := array[row+1][col+1]
	e := array[row+2][col]
	f := array[row+2][col+1]
	g := array[row+2][col+2]
	return a + b + c + d + e + f + g
}

func highest(vals []int) int {
	val := vals[0]
	for _, v := range vals[1:] {
		if v > val {
			val = v
		}
	}
	return val
}

func parse(a arr) [16]int {
	hourglasses := [16]int{}
	count := 0
	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			hourglasses[count] = hourglass(a, row, col)
			count++
		}
	}
	return hourglasses
}

func main() {
	a := input()
	hourglasses := parse(a)
	fmt.Print(highest(hourglasses[:]))
}
