package main

import "fmt"

func input() []int {
	val := make([]int, 3)
	for i := range val {
		fmt.Scanf("%d", &val[i])
	}
	return val
}

func main() {
	a := input()
	b := input()
	aPoints := 0
	bPoints := 0
	for i := range a {
		switch {
		case a[i] > b[i]:
			aPoints++
		case a[i] < b[i]:
			bPoints++
		}
	}
	fmt.Print(aPoints, bPoints)
}
