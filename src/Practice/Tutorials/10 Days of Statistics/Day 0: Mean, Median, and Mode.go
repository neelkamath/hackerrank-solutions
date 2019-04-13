package main

import (
	"fmt"
	"sort"
)

func mean(ints []int) float64 {
	sum := 0
	for _, v := range ints {
		sum += v
	}
	return float64(sum) / float64(len(ints))
}

func median(ints []int) float64 {
	sort.Ints(ints)
	var i int
	if len(ints)%2 == 0 {
		i = len(ints)/2 - 1
	} else {
		i = len(ints) / 2
	}
	return float64(ints[i]+ints[i+1]) / 2
}

func count(ints []int) map[int]int {
	m := make(map[int]int)
	for _, v := range ints {
		m[v] = m[v] + 1
	}
	return m
}

func largestValue(m map[int]int) int {
	var value int
	for _, v := range m {
		if v > value {
			value = v
		}
	}
	return value
}

func retainLargest(m map[int]int) {
	largest := largestValue(m)
	for k, v := range m {
		if v < largest {
			delete(m, k)
		}
	}
}

func keys(m map[int]int) []int {
	var ints []int
	for k := range m {
		ints = append(ints, k)
	}
	return ints
}

func min(m map[int]int) int {
	ints := keys(m)
	minimum := ints[0]
	for _, v := range ints[1:] {
		if v < minimum {
			minimum = v
		}
	}
	return minimum
}

func mode(ints []int) int {
	counts := count(ints)
	retainLargest(counts)
	return min(counts)
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	x := make([]int, n)
	for i := range x {
		fmt.Scanf("%d", &x[i])
	}
	fmt.Printf("%.1f\n%.1f\n%d\n", mean(x), median(x), mode(x))
}
