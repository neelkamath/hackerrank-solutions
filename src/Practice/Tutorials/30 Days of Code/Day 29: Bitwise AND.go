package main

import "fmt"

func set(n int) []int {
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = i + 1
	}
	return s
}

func and(ints []int) []int {
	var s []int
	for i := 0; i < len(ints)-1; i++ {
		for j := i + 1; j < len(ints); j++ {
			s = append(s, ints[i]&ints[j])
		}
	}
	return s
}

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for i := 0; i < t; i++ {
		var n, k int
		fmt.Scanf("%d %d", &n, &k)
		s := set(n)
		ands := and(s)
		var max int
		for _, v := range ands {
			if v > max && v < k {
				max = v
			}
		}
		fmt.Println(max)
	}
}
