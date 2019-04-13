package main

import "fmt"

func main() {
	var n, t int
	fmt.Scanf("%d %d", &n, &t)
	widths := make([]int, n)
	for i := range widths {
		fmt.Scanf("%d", &widths[i])
	}
	for i := 0; i < t; i++ {
		var i, j int
		fmt.Scanf("%d %d", &i, &j)
		smallest := 100000
		for _, width := range widths[i : j+1] {
			if width < smallest {
				smallest = width
			}
		}
		fmt.Println(smallest)
	}
}
