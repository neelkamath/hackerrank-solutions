package main

import "fmt"

func getMax(ints []int) int {
	max := ints[0]
	for _, v := range ints[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	var stack []int
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var queryType int
		fmt.Scan(&queryType)
		switch queryType {
		case 1:
			var x int
			fmt.Scan(&x)
			stack = append(stack, x)
		case 2:
			stack = stack[:len(stack)-1]
		case 3:
			fmt.Println(getMax(stack))
		}
	}
}
