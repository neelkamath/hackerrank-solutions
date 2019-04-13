package main

import (
	"fmt"
	"sort"
)

func gen(x, f []int) []int {
	var s []int
	for index, value := range x {
		for i := 0; i < f[index]; i++ {
			s = append(s, value)
		}
	}
	sort.Ints(s)
	return s
}

func median(x []int) float64 {
	length := len(x)
	halfLen := length / 2
	if length%2 == 0 {
		return float64(x[halfLen-1]+x[halfLen]) / 2
	}
	return float64(x[halfLen])
}

func calcQuartiles(s []int) (q1, q2, q3 float64) {
	q2 = median(s)
	length := len(s)
	halfLen := length / 2
	q1 = median(s[:halfLen])
	if length%2 == 0 {
		q3 = median(s[halfLen:])
	} else {
		q3 = median(s[halfLen+1:])
	}
	return
}

func calcIR(s []int) float64 {
	q1, _, q3 := calcQuartiles(s)
	return q3 - q1
}

func main() {
	var n int
	fmt.Scan(&n)
	x := make([]int, n)
	for i := range x {
		fmt.Scan(&x[i])
	}
	f := make([]int, n)
	for i := range f {
		fmt.Scan(&f[i])
	}
	s := gen(x, f)
	fmt.Printf("%.1f", calcIR(s))
}
