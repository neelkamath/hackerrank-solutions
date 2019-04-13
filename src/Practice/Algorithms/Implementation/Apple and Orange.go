package main

import "fmt"

func dist(size int) []int {
	distances := make([]int, size)
	for i := range distances {
		fmt.Scanf("%d", &distances[i])
	}
	return distances
}

func fallen(locations []int, location, begin, end int) int {
	fallen := 0
	for _, v := range locations {
		loc := location + v
		if loc >= begin && loc <= end {
			fallen++
		}
	}
	return fallen
}

func main() {
	var s, t, a, b, m, n int
	fmt.Scanf("%d %d\n%d %d\n%d %d", &s, &t, &a, &b, &m, &n)
	apples := dist(m)
	oranges := dist(n)

	fmt.Printf("%d\n%d", fallen(apples, a, s, t), fallen(oranges, b, s, t))
}
