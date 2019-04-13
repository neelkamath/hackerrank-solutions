package main

import "fmt"

func input(size int) []int {
	vals := make([]int, size)
	for i := range vals {
		fmt.Scanf("%d", &vals[i])
	}
	return vals
}

func main() {
	var b, n, m int
	fmt.Scanf("%d %d %d", &b, &n, &m)
	keyboard := input(n)
	drives := input(m)

	var bill int
	for _, k := range keyboard {
		for _, d := range drives {
			sum := k + d
			if sum > bill && sum <= b {
				bill = sum
			}
		}
	}
	if bill == 0 {
		fmt.Print(-1)
	} else {
		fmt.Print(bill)
	}
}
