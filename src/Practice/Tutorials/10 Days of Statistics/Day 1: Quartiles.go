package main

import (
	"fmt"
	"sort"
)

func median(x []int) int {
	length := len(x)
	halfLen := length / 2
	if length%2 == 0 {
		return (x[halfLen-1] + x[halfLen]) / 2
	}
	return x[halfLen]
}

func calcQuartiles(x []int) (q1, q2, q3 int) {
	q2 = median(x)
	length := len(x)
	halfLen := length / 2
	q1 = median(x[:halfLen])
	if length%2 == 0 {
		q3 = median(x[halfLen:])
	} else {
		q3 = median(x[halfLen+1:])
	}
	return
}

func main() {
	var n int
	fmt.Scan(&n)
	x := make([]int, n)
	for i := range x {
		fmt.Scan(&x[i])
	}
	sort.Ints(x)
	q1, q2, q3 := calcQuartiles(x)
	fmt.Printf("%d\n%d\n%d\n", q1, q2, q3)
}
