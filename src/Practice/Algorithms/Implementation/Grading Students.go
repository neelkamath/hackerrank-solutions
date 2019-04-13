package main

import "fmt"

func round(num, interval, diff int) int {
	val := 0
	for i := val; i < num+interval; i += interval {
		val = i
	}
	if val-num < diff {
		return val
	}
	return num
}

func main() {
	var n int
	fmt.Scanf("%d", &n)

	grades := make([]int, n)
	for i := range grades {
		fmt.Scanf("%d", &grades[i])
	}

	for _, grade := range grades {
		if grade < 38 {
			fmt.Println(grade)
		} else {
			fmt.Println(round(grade, 5, 3))
		}
	}
}
