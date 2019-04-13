package main

import "fmt"

func main() {
	var t int
	fmt.Scanf("%d", &t)
	for i := 0; i < t; i++ {
		var n, c, m int
		fmt.Scanf("%d %d %d", &n, &c, &m)
		wrappers := n / c
		chocolates := wrappers
		for wrappers >= m {
			bought := wrappers / m
			chocolates += bought
			wrappers = wrappers - bought*m + bought
		}
		fmt.Println(chocolates)
	}
}
